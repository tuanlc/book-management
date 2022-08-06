package models

import (
	"github.com/tuanlc/book-management/pkg/config"
	"github.com/tuanlc/book-management/pkg/types"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDBClient()
	db.AutoMigrate(&types.Book{})
}

func ListBooks() []types.Book {
	var books []types.Book
	db.Find(&books)

	return books
}

func DeleteBook(bookId int64) types.Book {
	var book types.Book

	db.Where("ID=?", bookId).Delete(&book)

	return book
}

func GetBook(bookId int64) *types.Book {
	var foundBook types.Book

	db.Where("ID=?", bookId).Find(&foundBook)

	return &foundBook
}

func CreateBook(b *types.Book) *types.Book {
	book := types.Book{Title: b.Title, Author: b.Author, Summary: b.Summary}

	db.Select("*").Create(&book)

	return &book
}

func UpdateBook(bookId int64, data *types.Book) *types.Book {
	var updateBook types.Book

	db.Model(&updateBook).Select("*").Where("ID=?", bookId).Updates(types.Book{Title: data.Title, Summary: data.Summary, Author: data.Author})

	return &updateBook
}
