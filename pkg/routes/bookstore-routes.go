package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/abh1SHAKE/go-crud/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router chi.Router) {
	router.Post("/book/", controllers.CreateBook)
	router.Get("/book/", controllers.GetBook)
	router.Get("/book/{bookId}", controllers.GetBookById)
	router.Put("/book/{bookId}", controllers.UpdateBook)
	router.Delete("/book/{bookId}", controllers.DeleteBook)
}