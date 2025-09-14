package middleware

import (
	"net/http"
)

func Cors(next http.Handler) http.Handler {
	handleAllCors := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Method", "POST, PUT,GET,DELETE,OPTIONS,PATCH") //allow post request
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")                      //allow content type for frotend
		w.Header().Set("Content-Type", "application/json")

		/// preflight request handle
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		next.ServeHTTP(w, r)
	})
	return http.HandlerFunc(handleAllCors)
}
