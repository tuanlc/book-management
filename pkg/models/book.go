package models

import (
	"time"

	"github.com/tuanlc/book-management/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	AuthorId  int       `json:"authorId"`
	Summary   string    `json:"summary"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func init() {
	config.Connect()
	db = config.GetDBClient()
	db.AutoMigrate(&Book{})
}

var books []Book

func ListBooks() []Book {
	var books []Book
	db.Find(&books)

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
	var foundBook Book

	db.Where("ID=?", bookId).Find(&foundBook)

	return &foundBook
}

func CreateBook(b *Book) *Book {
	book := Book{Title: b.Title, AuthorId: b.AuthorId, Summary: b.Summary}
	db.Select("ID", "Title", "Summary", "AuthorId", "CreatedAt", "UpdatedAt").Create(&book)

	return &book
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
