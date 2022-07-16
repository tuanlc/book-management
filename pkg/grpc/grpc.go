package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/tuanlc/book-management/pkg/grpc/book"
	"google.golang.org/grpc"
)

func Serve() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))

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
