package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("I'm middleware: I will print first.")

		start := time.Now();

		next.ServeHTTP(w, r)

		diff := time.Since(start)

		log.Println("I'm middleware: I will print last.") 
		log.Println(r.Method, r.URL.Path, diff)
	})
}	