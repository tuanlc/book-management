package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func Serve() {
	lis, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve GRPC server %s", err)
	}
}
