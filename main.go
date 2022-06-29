package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Author  *Author `json:"author"`
	Summary string  `json:"summary"`
}

type Author struct {
	name string `json:"name"`
}

var books []Book

func listBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(10000000000))

	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
}

func main() {
	r := mux.NewRouter()

	books = append(books, Book{
		ID:    "1",
		Title: "Sapiens",
		Author: &Author{
			name: "Yuval Noah Harari",
		},
	})

	books = append(books, Book{
		ID:    "2",
		Title: "Khi Loi Thuoc Ve Nhung Vi Sao",
		Author: &Author{
			name: "John Green",
		},
	})

	r.HandleFunc("/books", listBooks).Method("GET")
	r.HandleFunc("/books/[id]", getBook).Method("GET")
	r.HandleFunc("/books", createBook).method("POST")
	r.HandleFunc("/books/[id]", updateBook).method("PATCH")
	r.HandleFunc("/books/[id]", deleteBook).method("DELETE")

	fmt.Printf("Starting web server at port 8000\n")

	log.Fatal(http.ListenAndServe(":8000", r))
}
