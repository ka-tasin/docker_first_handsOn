package main

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/product"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Go server is running")
	}))

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProduct))
	mux.Handle("POST /create-product", http.HandlerFunc(handlers.CreateProduct))

	fmt.Println("Server is running on port: 8080")

	globalRouter := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(":8080", globalRouter)

	if err != nil {
		fmt.Println("Error Starting the Server", err)
	}
}

func init() {
	prod1 := product.Product{
		ID:          1,
		Title:       "Orane",
		Description: "Best orange ever",
		Price:       20,
		ImageUrl:    "https://www.quanta.org/orange/",
	}

	prod2 := product.Product{
		ID:          2,
		Title:       "Apple",
		Description: "Best apple ever",
		Price:       20,
		ImageUrl:    "https://www.quanta.org/orange/",
	}

	prod3 := product.Product{
		ID:          3,
		Title:       "Banana",
		Description: "Best banana ever",
		Price:       20,
		ImageUrl:    "https://www.quanta.org/orange/",
	}

	product.ProductList = append(product.ProductList, prod1)
	product.ProductList = append(product.ProductList, prod2)
	product.ProductList = append(product.ProductList, prod3)
}



