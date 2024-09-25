package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	grpc_server "main/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	grpc_server.UnimplementedHealthCheckServiceServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) HealthCheck(ctx context.Context, req *grpc_server.HealthCheckRequest) (*grpc_server.HealthCheckResponse, error) {
	return &grpc_server.HealthCheckResponse{
		Message: "HealthCheck, OK",
	}, nil
}

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	grpc_server.RegisterHealthCheckServiceServer(s, NewServer())

	reflection.Register(s)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
