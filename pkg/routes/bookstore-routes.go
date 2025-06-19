package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/abh1SHAKE/go-crud/pkg/controllers"
)

func RegisterBookStoreRoutes(router chi.Router) {
	router.Post("/books/", controllers.CreateBook)
	router.Get("/books/", controllers.GetBooks)
	router.Get("/books/{id}", controllers.GetBookById)
	router.Put("/books/{id}", controllers.UpdateBook)
	router.Delete("/books/{id}", controllers.DeleteBook)
}