package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tuanlc/book-management/pkg/models"
	"github.com/tuanlc/book-management/pkg/types"
	"github.com/tuanlc/book-management/pkg/utils"
)

func ListBooks(w http.ResponseWriter, r *http.Request) {
	books := models.ListBooks()

	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(books)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	ID, err := strconv.ParseInt(params["id"], 0, 0)

	if err != nil {
		fmt.Printf("Error while parsing ID")
	}

	book := models.DeleteBook(ID)

	res, _ := json.Marshal(book)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	ID, err := strconv.ParseInt(params["id"], 0, 0)

	if err != nil {
		fmt.Printf("Error while parsing ID")
	}

	book := models.GetBook(ID)

	res, _ := json.Marshal(book)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	book := &types.Book{}
	utils.ParseBody(r, book)

	createdBook := models.CreateBook(book)

	res, _ := json.Marshal(createdBook)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	book := &types.Book{}
	utils.ParseBody(r, book)

	ID, err := strconv.ParseInt(params["id"], 0, 0)

	if err != nil {
		fmt.Printf("Error while parsing ID")
	}

	updatedBook := models.UpdateBook(ID, book)

	res, _ := json.Marshal(updatedBook)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
