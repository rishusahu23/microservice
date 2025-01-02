package main

import (
	"context"
	"fmt"
	"github.com/rishu/microservice/config"
	userPb "github.com/rishu/microservice/gen/api/user"
	"github.com/rishu/microservice/pkg/db/mongo"
	redis2 "github.com/rishu/microservice/pkg/in_memory_store/redis"
	"github.com/rishu/microservice/user/wire"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"log"
	"net"
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

	go func() {
		startGrpcServer(ctx, grpcL, conf)
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

func main() {
	ctx := context.Background()
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}

	// Start combined gRPC and HTTP server
	startCombinedServer(ctx, conf)

	// Block main goroutine indefinitely (this will keep the servers running)
	select {}
}
