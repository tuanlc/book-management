package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/tuanlc/book-management/pkg/grpc/book"
	"google.golang.org/grpc"
)

func Serve() {
	fmt.Printf("Starting gRPC server!")

	hostport := "localhost:9090"
	lis, err := net.Listen("tcp", hostport)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := book.Server{}

	grpcServer := grpc.NewServer()

	book.RegisterBookServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve GRPC server %s", err)
	}
}
