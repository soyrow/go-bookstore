package routes

import (
	"github.com/gorilla/mux"
	"github.com/soyrow/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Route) {
	router.HandlerFunc("./book/", controllers.CreateBook).Methods("POST")
	router.HandlerFunc("./book/", controllers.GetBook).Methods("GET")
	router.HandlerFunc("./book/", controllers.GetBookById).Methods("GET")
	router.HandlerFunc("./book/", controllers.UpdateBook).Methods("PUT")
	router.HandlerFunc("./book/", controllers.DeleteBook).Methods("DELETE")
}
