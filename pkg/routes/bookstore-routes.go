package routes

import (
	"go-bookstore/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBooKStoreRoutes = func(router *mux.NewRouter) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("Delete")
}
