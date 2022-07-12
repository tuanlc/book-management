package main

import (
	"github.com/tuanlc/book-management/pkg/grpc"
	"github.com/tuanlc/book-management/pkg/rest"
)

func main() {
	rest.Serve()
	grpc.Serve()
}
