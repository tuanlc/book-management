package routes

import (
	"github.com/gorilla/mux"
	"github.com/tuanlc/book-management/pkg/rest/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.ListBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PATCH")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
}
