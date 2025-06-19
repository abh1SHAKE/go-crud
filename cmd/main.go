package main 

import (
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/abh1SHAKE/go-crud/pkg/routes"
)

func main() {
	r := chi.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	log.Println("Server started at http://localhost:7744")
	log.Fatal(http.ListenAndServe("localhost:7744", r))
}