package global_router

import "net/http"

func GlobalRouter(mux *http.ServeMux) http.HandlerFunc {
	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Controll-Allow-Origin", "*")
		w.Header().Set("Access-Controll-Allow-Origin", "GET, POST, PUT PATCH DELETE OPTIONS")
		w.Header().Set("Access-Controll-Allow-Origin", "Content-Type, tata")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodOptions {
			w.WriteHeader(200)
			return
		}

		mux.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handleAllReq)
}