package handlers

import (
	"ecommerce/product"
	"ecommerce/utils"
	"net/http"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Please give a get request.", 400)
		return
	}

	utils.SendData(w, product.ProductList, 200)

}
