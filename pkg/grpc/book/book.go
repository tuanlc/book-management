package book

import (
	"context"
	"log"

	"github.com/tuanlc/book-management/pkg/models"
	"github.com/tuanlc/book-management/pkg/types"
)

type Server struct {
	UnimplementedBookServiceServer
}

func (s *Server) GetBook(ctx context.Context, in *GetBookRequest) (*GetBookResponse, error) {
	book := models.GetBook(in.Id)

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

func (s *Server) ListBooks(ctx context.Context, in *ListBookRequest) (*ListBookResponse, error) {
	books := models.ListBooks()

	responseBooks := make([]*Book, len(books))

	for i := 0; i < len(books); i++ {
		book := books[i]
		responseBooks[i] = &Book{Id: book.ID, Title: book.Title, Summary: book.Summary, Author: book.Author, CreatedAt: book.CreatedAt.String(), UpdatedAt: book.UpdatedAt.String()}
	}

	return &ListBookResponse{
		Books: responseBooks,
	}, nil
}

func (s *Server) CreateBook(ctx context.Context, in *CreateBookRequest) (*CreateBookResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()

	book := models.CreateBook(&types.CreateBookPayload{Title: in.Title, Author: in.Author, Summary: *in.Summary})

	return &CreateBookResponse{
		Book: &Book{Id: book.ID, Title: book.Title, Summary: book.Summary, Author: book.Author, CreatedAt: book.CreatedAt.String(), UpdatedAt: book.UpdatedAt.String()},
	}, nil
}
