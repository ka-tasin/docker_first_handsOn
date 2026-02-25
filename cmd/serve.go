package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

	manager.Use(middleware.Hudai, middleware.Logger)

	mux := http.NewServeMux()

	initRoutes(mux, manager)

	fmt.Println("Server is running on port: 8080")

	globalRouter := middleware.CorsWithPreflight(mux)

	err := http.ListenAndServe(":8080", globalRouter)

	if err != nil {
		fmt.Println("Error Starting the Server", err)
	}
}
