package rest

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tuanlc/book-management/pkg/rest/routes"
)

func Serve() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
