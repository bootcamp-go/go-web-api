package main

import (
	"app/cmd/handlers"
	"app/internal/products/storage"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// env
	// ...

	// dependencies
	// -> storage
	stBase := storage.NewStorageJsonFile(os.Getenv("STORAGE_JSON_FILE_PATH"))
	st := storage.NewStorageProductDefault(stBase)
	// -> controller
	ct := handlers.NewControllerProducts(st)
	
	// server
	rt := chi.NewRouter()
	// -> middlewares
	rt.Use(middleware.Recoverer)
	rt.Use(middleware.Logger)
	// -> handlers
	rt.Route("/api/v1", func(rt chi.Router) {
		// -> products
		rt.Route("/products", func(rt chi.Router) {
			// -> GET /products
			rt.Get("/", ct.GetAll())
			// -> GET /products/{id}
			rt.Get("/{id}", ct.GetByID())
			// -> GET /products/search (query params)
			rt.Get("/search", ct.Search())
		})
	})
	// run
	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}