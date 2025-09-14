package routes

import (
	"net/http"

	"github.com/saadahmedbd/Treestore/Rest/Handler/SearchHandler"
	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func Search_ProductRoute(mux *http.ServeMux) {
	mux.Handle("GET /search", middleware.Chain(http.HandlerFunc(searchhandler.Search_Product),
		middleware.Logger,
		middleware.Cors,
	))
	mux.Handle("GET /product/bestselling", middleware.Chain(http.HandlerFunc(searchhandler.BestSellingProduct),
		middleware.Logger,
		middleware.Cors,
	))
	mux.Handle("GET /product/topproduct", middleware.Chain(http.HandlerFunc(searchhandler.TopProduct),
		middleware.Logger,
		middleware.Cors,
	))
}
