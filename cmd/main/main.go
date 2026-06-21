package main

import (
	"log"
	"net/http"

	"github.com/emious/go-book-management-api/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Println("Server started on: http://localhost:9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
