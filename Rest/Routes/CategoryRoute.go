package routes

import (
	"net/http"

	categoryHandler "github.com/saadahmedbd/Treestore/Rest/Handler/CategoryHandler"
	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func CategoryRoute(mux *http.ServeMux) {
	mux.Handle("GET /getcategory", middleware.Chain(http.HandlerFunc(categoryHandler.GetCategory),
		middleware.Logger,
	))
	mux.Handle("POST /createcategory", middleware.Chain(http.HandlerFunc(categoryHandler.CreateCategory),
		middleware.Logger,
	))
	mux.Handle("GET /getcategory/{categoryId}", middleware.Chain(http.HandlerFunc(categoryHandler.GetCategoryById),
		middleware.Logger,
	))
	mux.Handle("PUT /updatecategory/{categoryId}", middleware.Chain(http.HandlerFunc(categoryHandler.UpdateCategory),
		middleware.Logger,
	))
	mux.Handle("DELETE /deletecategory/{categoryId}", middleware.Chain(http.HandlerFunc(categoryHandler.DeleteCategory),
		middleware.Logger,
	))
}
