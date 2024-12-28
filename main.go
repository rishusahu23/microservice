package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rishu/microservice/config"
	userPb "github.com/rishu/microservice/gen/api/user"
	"github.com/rishu/microservice/pkg/db/mongo"
	redis2 "github.com/rishu/microservice/pkg/in_memory_store/redis"
	"github.com/rishu/microservice/user/wire"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
	"net"
	"net/http"
)

// Combined gRPC and HTTP server using cmux
func startCombinedServer(ctx context.Context, conf *config.Config) {
	// Create a listener for the shared port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", conf.Server.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen on port %v: %v", conf.Server.GrpcPort, err)
	}

	// Create a cmux instance
	m := cmux.New(lis)

	// Match connections for gRPC
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))

	// Match connections for HTTP (gRPC-Gateway and custom HTTP)
	httpL := m.Match(cmux.Any())

	// Start gRPC server
	go func() {
		startGrpcServer(ctx, grpcL, conf)
	}()

	// Start HTTP server
	go func() {
		startGrpcHttpServer(ctx, httpL, conf)
	}()

	// Start cmux serving
	log.Printf("Starting combined gRPC and HTTP server on :%v", conf.Server.GrpcPort)
	if err := m.Serve(); err != nil {
		log.Fatalf("cmux server error: %v", err)
	}
}

// gRPC server setup
func startGrpcServer(ctx context.Context, lis net.Listener, conf *config.Config) {
	opts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(1024 * 1024 * 100), // Max receive message size (100 MB)
		grpc.MaxSendMsgSize(1024 * 1024 * 100), // Max send message size (100 MB)
	}

	s := grpc.NewServer(opts...)
	mongoClient := mongo.GetMongoClient(ctx, conf)
	redisClient := redis2.GetRedisClient(conf)
	userPb.RegisterUserServiceServer(s, wire.InitialiseUserService(conf, mongoClient, redisClient))

	log.Printf("gRPC server running")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}

// gRPC-Gateway server setup
func startGrpcHttpServer(ctx context.Context, lis net.Listener, conf *config.Config) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := userPb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%v", conf.Server.GrpcPort), opts)
	if err != nil {
		log.Fatalf("failed to register service handler: %v", err)
	}

	httpServer := &http.Server{
		Handler: mux, // Use the mux directly
	}

	log.Printf("HTTP server running")
	if err := httpServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve HTTP server: %v", err)
	}
}

// HTTP server setup
func startHttpServer(conf *config.Config) {
	// Use gorilla/mux for routing HTTP requests
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from HTTP API"))
	}).Methods("GET")

	// Start HTTP server on port 8150
	log.Printf("Starting HTTP server on :%v", conf.Server.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", conf.Server.Port), router); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}

func main() {
	ctx := context.Background()
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}
	// Start gRPC server in a separate goroutine
	//go startGrpcServer(ctx, conf)

	// Start HTTP server in the main goroutine
	go startHttpServer(conf)

	// Optionally, start gRPC-over-HTTP (requires grpc-gateway setup)
	//go startGrpcHttpServer(conf)

	// Start combined gRPC and HTTP server
	startCombinedServer(ctx, conf)

	// Block main goroutine indefinitely (this will keep the servers running)
	select {}
}

func statusToHTTPCode(statusCode codes.Code) int {
	// Map custom gRPC status code to HTTP status code
	switch statusCode {
	case codes.OK:
		return http.StatusOK
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.NotFound:
		return http.StatusNotFound
	case codes.Internal:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

// Response structure to match gRPC response format
type GrpcResponse struct {
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	User interface{} `json:"user"` // Assuming `user` is dynamic
}

// Middleware for extracting status code from the gRPC response
func grpcGatewayMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Capture the original response using a response recorder
		rec := &responseRecorder{ResponseWriter: w, body: make([]byte, 0)}

		// Call the next handler (gRPC-Gateway handler)
		next.ServeHTTP(rec, r)

		// Parse the captured response body
		var grpcResponse GrpcResponse
		if err := json.Unmarshal(rec.body, &grpcResponse); err != nil {
			log.Printf("Failed to parse response body: %v", err)
			// Fall back to internal server error if parsing fails
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Map gRPC status code to HTTP status code
		httpStatusCode := statusToHTTPCode(codes.Code(grpcResponse.Status.Code))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpStatusCode) // Set the correct HTTP status code
		_, err := w.Write(rec.body)   // Write the original body
		if err != nil {
			log.Printf("Failed to write response: %v", err)
		}
	})
}

// Custom response recorder to capture the response body
type responseRecorder struct {
	http.ResponseWriter
	body []byte
}

func (r *responseRecorder) Write(p []byte) (n int, err error) {
	// Append response body to the recorder's body field
	r.body = append(r.body, p...)
	return len(p), nil
}
