package routes

import (
	"net/http"

	producthandler "github.com/saadahmedbd/Treestore/Handler/ProductHandler"
	middleware "github.com/saadahmedbd/Treestore/Middleware"
)

func ProductRoute(mux *http.ServeMux) {
	mux.Handle("GET /getproduct", middleware.Chain(http.HandlerFunc(producthandler.GetProduct),
		middleware.Logger,
	))
	mux.Handle("POST /createproduct", middleware.Chain(http.HandlerFunc(producthandler.CreateProduct),
		middleware.Logger,
	))
	mux.Handle("GET /getproduct/{productId}", middleware.Chain(http.HandlerFunc(producthandler.GetProductById),
		middleware.Logger,
	))
	mux.Handle("PUT /updateproduct/{productId}", middleware.Chain(http.HandlerFunc(producthandler.UpdateProduct),
		middleware.Logger,
	))
	mux.Handle("DELETE /deleteproduct/{productId}", middleware.Chain(http.HandlerFunc(producthandler.DeleteProduct),
		middleware.Logger,
	))
}
