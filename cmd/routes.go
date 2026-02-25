package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /route", 
		manager.With(
			http.HandlerFunc(handlers.Test),
			middleware.Arekta,
		),
	)

	mux.Handle("GET /products", middleware.Logger(http.HandlerFunc(handlers.GetProduct)))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{productId}", http.HandlerFunc(handlers.GetProductById))

	mux.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Go server is running")
	}))
}