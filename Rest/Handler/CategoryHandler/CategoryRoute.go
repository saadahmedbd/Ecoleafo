package categoryHandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *CategoryHandler) CategoryRoute(mux *http.ServeMux) {
	//public route no authentication required
	mux.Handle("GET /api/categories", http.HandlerFunc(h.GetAllCategories))
	mux.Handle("GET /api/categories/get", http.HandlerFunc(h.GetCategoryByID))
	mux.Handle("GET /api/categories/slug", http.HandlerFunc(h.GetCategoryBySlug))
	mux.Handle("GET /api/categories/root", http.HandlerFunc(h.GetRootCategories))
	mux.Handle("GET /api/categories/subcategories", http.HandlerFunc(h.GetSubcategories))
	mux.Handle("GET /api/categories/featured", http.HandlerFunc(h.GetFeaturedCategories))
	mux.Handle("GET /api/categories/tree", http.HandlerFunc(h.GetCategoryTree))
	mux.Handle("GET /api/categories/breadcrumb", http.HandlerFunc(h.GetBreadcrumb))

	//only admin routes (authentication required)
	mux.Handle("POST /api/categories/create", middleware.Chain(http.HandlerFunc(h.CreateCategory),
		middleware.Logger,
		middleware.AuthenticateJWT,
		middleware.Cors,
	))

	mux.Handle("PUT /api/categories/update", middleware.Chain(http.HandlerFunc(h.UpdateCategory),
		middleware.Logger,
		middleware.AuthenticateJWT,
		middleware.Cors,
	))
	mux.Handle("DELETE /categories/delete", middleware.Chain(http.HandlerFunc(h.DeleteCategory),
		middleware.Logger,
		middleware.AuthenticateJWT,
		middleware.Cors,
	))
	mux.Handle("POST /api/categories/reorder", middleware.Chain(http.HandlerFunc(h.ReorderCategories),
		middleware.Logger,
		middleware.AuthenticateJWT,
		middleware.Cors,
	))
}
