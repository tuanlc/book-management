package book

import (
	"context"
	"log"
)

type Server struct {
	UnimplementedBookServiceServer
}

func (s *Server) GetBook(ctx context.Context, in *GetBookRequest) (*Book, error) {
	log.Printf("Receive request from client %v\n", in.Id)
	return &Book{Id: 1, Title: "Test", Summary: "Test summary", Author: "Tao", CreatedAt: "20:00", UpdatedAt: "21:00"}, nil
}
