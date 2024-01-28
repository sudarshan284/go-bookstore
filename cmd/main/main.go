package main

import (
	"log"
	"net/http"

	"pkg/routes"

	"github.com/gorilla/mux"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.HandleFunc("/", r)
	log.Fatal(http.ListenAndServe(":9000", r))
}
