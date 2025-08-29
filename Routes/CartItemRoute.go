package routes

import (
	"net/http"

	categoryHandler "github.com/saadahmedbd/Treestore/Handler/CategoryHandler"
	middleware "github.com/saadahmedbd/Treestore/Middleware"
)

func CartItemRoute(mux *http.ServeMux) {
	mux.Handle("GET /getcartitem", middleware.Chain(http.HandlerFunc(categoryHandler.GetCategory),
		middleware.Logger,
	))
	mux.Handle("POST /createcartitem", middleware.Chain(http.HandlerFunc(categoryHandler.CreateCategory),
		middleware.Logger,
	))
	mux.Handle("GET /getcartitem/{cartitemId}", middleware.Chain(http.HandlerFunc(categoryHandler.GetCategoryById),
		middleware.Logger,
	))
	mux.Handle("PUT /updatecartitem/{cartitemId}", middleware.Chain(http.HandlerFunc(categoryHandler.UpdateCategory),
		middleware.Logger,
	))
	mux.Handle("DELETE /deletecartitem/{cartitemId}", middleware.Chain(http.HandlerFunc(categoryHandler.DeleteCategory),
		middleware.Logger,
	))
}
