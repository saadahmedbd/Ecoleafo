package categoryHandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *Handler) CategoryRoute(mux *http.ServeMux) {
	mux.Handle("GET /getcategory", middleware.Chain(http.HandlerFunc(h.GetCategory),
		middleware.Logger,
	))
	mux.Handle("POST /createcategory", middleware.Chain(http.HandlerFunc(h.CreateCategory),
		middleware.Logger,
		middleware.AuthenticateJWT,
	))
	mux.Handle("GET /getcategory/{categoryId}", middleware.Chain(http.HandlerFunc(h.GetCategoryById),
		middleware.Logger,
	))
	mux.Handle("PUT /updatecategory/{categoryId}", middleware.Chain(http.HandlerFunc(h.UpdateCategory),
		middleware.Logger,
	))
	mux.Handle("DELETE /deletecategory/{categoryId}", middleware.Chain(http.HandlerFunc(h.DeleteCategory),
		middleware.Logger,
	))
}
