package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	config "github.com/rishu/microservice/config"
	userPb "github.com/rishu/microservice/gen/api/user"
	"github.com/rishu/microservice/pkg/db/mongo"
	"github.com/rishu/microservice/user/wire"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

// gRPC server setup
func startGrpcServer(ctx context.Context, conf *config.Config) {
	// Listen on the gRPC port (9090)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", conf.Server.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen on port %v: %v", conf.Server.GrpcPort, err)
	}

	// Create the gRPC server
	s := grpc.NewServer()
	mongoClient := mongo.GetMongoClient(ctx, conf)
	userPb.RegisterUserServiceServer(s, wire.InitialiseUserService(conf, mongoClient))

	log.Printf("Starting gRPC server on :%v", conf.Server.GrpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
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

// gRPC-over-HTTP (optional, for gRPC-gateway)
func startGrpcHttpServer(conf *config.Config) {
	// Create a reverse proxy for gRPC-Gateway
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()} // gRPC dial options for the gateway

	// Register the gRPC service to the HTTP reverse proxy
	err := userPb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%v", conf.Server.GrpcPort), opts)
	if err != nil {
		log.Fatalf("failed to register service handler: %v", err)
	}

	// Start the HTTP server (HTTP/REST interface)
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%v", conf.Server.GrpcHttpPort), // HTTP port for gRPC over HTTP
		Handler: mux,                                          // HTTP server uses the reverse proxy
	}

	log.Printf("Starting gRPC over HTTP server on :9091")
	if err := httpServer.ListenAndServe(); err != nil {
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
	go startGrpcServer(ctx, conf)

	// Start HTTP server in the main goroutine
	go startHttpServer(conf)

	// Optionally, start gRPC-over-HTTP (requires grpc-gateway setup)
	go startGrpcHttpServer(conf)

	// Block main goroutine indefinitely (this will keep the servers running)
	select {}
}
