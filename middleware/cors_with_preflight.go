package middleware

import "net/http"

func CorsWithPreflight(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Controll-Allow-Origin", "*")
		w.Header().Set("Access-Controll-Allow-Origin", "GET, POST, PUT PATCH DELETE OPTIONS")
		w.Header().Set("Access-Controll-Allow-Origin", "Content-Type, tata")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodOptions {
			w.WriteHeader(200)
			return
		}

		next.ServeHTTP(w, r)
	})
}