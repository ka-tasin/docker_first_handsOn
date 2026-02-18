package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"log"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Go server is running")
	}))

	middle := func(w http.ResponseWriter, r *http.Request) {
		log.Println("I am route, I will print in middle")
	}

	mux.Handle("GET /route", middleware.Logger(http.HandlerFunc(middle)))

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProduct))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))

	fmt.Println("Server is running on port: 8080")

	globalRouter := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(":8080", globalRouter)

	if err != nil {
		fmt.Println("Error Starting the Server", err)
	}
}