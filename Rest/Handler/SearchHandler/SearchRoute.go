package searchhandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *Handler) Search_ProductRoute(mux *http.ServeMux) {
	mux.Handle("GET /search", middleware.Chain(http.HandlerFunc(h.Search_Product),
		middleware.Logger,
		middleware.Cors,
	))
	mux.Handle("GET /product/bestselling", middleware.Chain(http.HandlerFunc(h.BestSellingProduct),
		middleware.Logger,
		middleware.Cors,
	))
	mux.Handle("GET /product/topproduct", middleware.Chain(http.HandlerFunc(h.TopProduct),
		middleware.Logger,
		middleware.Cors,
	))
}
