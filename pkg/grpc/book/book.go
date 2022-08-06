package book

import (
	"context"
	"fmt"
	"log"

	"github.com/tuanlc/book-management/pkg/models"
)

type Server struct {
	UnimplementedBookServiceServer
}

func (s *Server) GetBook(ctx context.Context, in *GetBookRequest) (*GetBookResponse, error) {
	log.Printf("Receive request from client %v\n", in.Id)

	book := models.GetBook(in.Id)

	fmt.Printf("found book %v", &GetBookResponse{
		Error: &ResponseError{
			Code:    "404",
			Message: "The book is not found",
		},
	})

	if book.ID == 0 {
		return &GetBookResponse{
			Error: &ResponseError{
				Code:    "404",
				Message: "The book is not found",
			},
		}, nil
	}

	return &GetBookResponse{
		Book: &Book{Id: book.ID, Title: book.Title, Summary: book.Summary, Author: book.Author, CreatedAt: book.CreatedAt.String(), UpdatedAt: book.UpdatedAt.String()},
	}, nil
}
