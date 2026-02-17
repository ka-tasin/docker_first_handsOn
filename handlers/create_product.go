package handlers

import (
	"ecommerce/product"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Please Provide Post Request.", 400)
		return
	}

	var newProduct product.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please provide valid JSON", 400)
		return
	}

	newProduct.ID = len(product.ProductList) + 1
	product.ProductList = append(product.ProductList, newProduct)

	utils.SendData(w, newProduct, 201)

}