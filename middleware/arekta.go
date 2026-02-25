package middleware

import (
	"fmt"
	"net/http"
)

func Arekta(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("I am arekta middleware")
		next.ServeHTTP(w, r)
	}) 
}