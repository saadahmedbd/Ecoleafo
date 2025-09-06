package routes

import (
	"net/http"

	searchhandler "github.com/saadahmedbd/Treestore/Handler/SearchHandler"
	middleware "github.com/saadahmedbd/Treestore/Middleware"
)

func Search_ProductRoute(mux *http.ServeMux) {
	mux.Handle("GET /search", middleware.Chain(http.HandlerFunc(searchhandler.Search_Product),
		middleware.Logger,
		middleware.Cors,
	))
}
