package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`
}

var productList []Product

func getProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Please give a get request.", 400)
		return
	}

	sendData(w, productList, 200)

}

func createProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Please Provide Post Request.", 400)
		return
	}

	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please provide valid JSON", 400)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	sendData(w, newProduct, 201)

}


func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Go server is running")
	}))

	mux.Handle("GET /products", http.HandlerFunc(getProduct))
	mux.Handle("POST /create-product", http.HandlerFunc(createProduct))

	fmt.Println("Server is running on port: 8080")

	globalRouter := globalRouter(mux)

	err := http.ListenAndServe(":8080", globalRouter)

	if err != nil {
		fmt.Println("Error Starting the Server", err)
	}
}

func init() {
	prod1 := Product{
		ID:          1,
		Title:       "Orane",
		Description: "Best orange ever",
		Price:       20,
		ImageUrl:    "https://www.quanta.org/orange/",
	}

	prod2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Best apple ever",
		Price:       20,
		ImageUrl:    "https://www.quanta.org/orange/",
	}

	prod3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Best banana ever",
		Price:       20,
		ImageUrl:    "https://www.quanta.org/orange/",
	}

	productList = append(productList, prod1)
	productList = append(productList, prod2)
	productList = append(productList, prod3)
}


func globalRouter(mux *http.ServeMux) http.HandlerFunc {
	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Controll-Allow-Origin", "*")
		w.Header().Set("Access-Controll-Allow-Origin", "GET, POST, PUT PATCH DELETE OPTIONS")
		w.Header().Set("Access-Controll-Allow-Origin", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		if(r.Method == http.MethodOptions) {
			w.WriteHeader(200)
			return
		}

		mux.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handleAllReq)
}
