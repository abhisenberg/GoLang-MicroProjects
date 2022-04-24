package routes

import (
	"github.com/abhisenberg/bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

/**
Defining the routes here (URL path and type of request)
and matching them to the suitable handler functions.
Params: router
**/
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
