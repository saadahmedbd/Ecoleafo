package routehandler

import "net/http"

func GlobalHandler(mux *http.ServeMux) http.Handler {
	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		//handle cors error
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Method", "POST, PUT,GET,DELETE,OPTIONS,PATCH") //allow post request
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")                      //allow content type for frotend
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		mux.ServeHTTP(w, r)
	}
	return http.HandlerFunc(handleAllReq)
}
