package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

	manager.Use(middleware.Hudai, middleware.Logger)

	mux := http.NewServeMux()

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

	fmt.Println("Server is running on port: 8080")

	globalRouter := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(":8080", globalRouter)

	if err != nil {
		fmt.Println("Error Starting the Server", err)
	}
}
