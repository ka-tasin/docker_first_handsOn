package middleware

import (
	"fmt"
	"net/http"
)

func Hudai(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("I am hudai middleware")
		next.ServeHTTP(w, r)
	}) 
}