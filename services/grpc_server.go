package services

import (
	"log"
	"net"
	"todo-api/routes"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGRPCServer(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	routes.RegisterGRPCRoutes(s)

	// Bật gRPC Reflection
	reflection.Register(s)

	log.Printf("gRPC server listening at %v", address)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
