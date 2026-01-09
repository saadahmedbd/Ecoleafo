package searchhandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *Handler) Search_ProductRoute(mux *http.ServeMux) {
	mux.Handle("GET /api/search", middleware.Chain(http.HandlerFunc(h.Search_Product),
		middleware.Logger,
		middleware.Cors,
	))
	mux.Handle("GET /api/search/autocomplete", middleware.Chain(http.HandlerFunc(h.Autocomplete),
		middleware.Logger,
		middleware.Cors,
	))
	mux.Handle("GET /api/product/bestselling", middleware.Chain(http.HandlerFunc(h.BestSellingProduct),
		middleware.Logger,
		middleware.Cors,
	))
	mux.Handle("GET /api/product/topproduct", middleware.Chain(http.HandlerFunc(h.TopProduct),
		middleware.Logger,
		middleware.Cors,
	))
}
