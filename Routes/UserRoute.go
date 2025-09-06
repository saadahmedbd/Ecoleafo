package routes

import (
	"net/http"

	userhandler "github.com/saadahmedbd/Treestore/Handler/UserHandler"
	middleware "github.com/saadahmedbd/Treestore/Middleware"
)

func UserRoute(mux *http.ServeMux) {
	mux.HandleFunc("/home", userhandler.Home)

	mux.Handle("GET /getseller", middleware.Chain(http.HandlerFunc(userhandler.GetSeller),
		middleware.Cors,
		middleware.Logger,
	)) //Getseller route

	mux.Handle("POST /createseller", middleware.Chain(http.HandlerFunc(userhandler.CreateSeller),
		middleware.Cors,
		middleware.Logger,
	)) // create seller route
	mux.Handle("GET /getseller/{sellerId}", middleware.Chain(http.HandlerFunc(userhandler.GetSellerById),
		middleware.Cors,
		middleware.Logger,
	)) // get seller by id route
	mux.Handle("PUT /updateseller/{sellerId}", middleware.Chain(http.HandlerFunc(userhandler.UpdateSeller),
		middleware.Cors,
		middleware.Logger,
	)) //update seller route
	mux.Handle("DELETE /deleteseller/{sellerId}", middleware.Chain(http.HandlerFunc(userhandler.DeleteSeller),
		middleware.Cors,
		middleware.Logger,
	)) //update seller
}
