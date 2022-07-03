package models

import (
	"github.com/tuanlc/book-management/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	AuthorId int    `json:"authorId"`
	Summary  string `json:"summary"`
}

func init() {
	config.Connect()
	db = config.GetDBClient()
	db.AutoMigrate(&Book{})
}

var books []Book

func ListBooks() []Book {
	return books
}

func DeleteBook(bookId int64) *[]Book {
	for index, item := range books {
		if item.ID == bookId {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	return &books
}

func GetBook(bookId int64) *Book {
	var foundBook *Book

	for _, item := range books {
		if item.ID == bookId {
			foundBook = &item
			break
		}
	}

	return foundBook
}

func CreateBook(book *Book) *Book {
	books = append(books, *book)

	return book
}

func UpdateBook(bookId int64, data *Book) *Book {
	var updateBook *Book

	for index, item := range books {
		if item.ID == bookId {
			books[index] = *data
			break
		}
	}

	return updateBook
}
