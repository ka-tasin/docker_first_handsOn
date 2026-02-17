package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Please give a get request.", 400)
		return
	}

	utils.SendData(w, database.ProductList, 200)

}
