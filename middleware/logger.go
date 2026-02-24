package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		start := time.Now()

		next.ServeHTTP(w, r)

		fmt.Println(r.Method, r.URL.Path, time.Since(start))
	})
}