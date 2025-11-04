package adminmangementhandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *ProductHandler) RegisterProductRoutes(mux *http.ServeMux) {
	//this api already have in product routes (api/product)
	mux.Handle("GET /api/admin/products", middleware.Chain(http.HandlerFunc(h.GetAllProducts),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	mux.Handle("GET /api/products/pending", middleware.Chain(http.HandlerFunc(h.GetPendingApprovals),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	mux.Handle("GET /api/products/search", middleware.Chain(http.HandlerFunc(h.SearchProducts),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	mux.Handle("GET /api/products/stats", middleware.Chain(http.HandlerFunc(h.GetProductStats),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	mux.Handle("GET /api/products/get", middleware.Chain(http.HandlerFunc(h.GetProductByID),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	mux.Handle("POST /api/products/approve", middleware.Chain(http.HandlerFunc(h.ApproveProduct),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	mux.Handle("POST /api/products/reject", middleware.Chain(http.HandlerFunc(h.RejectProduct),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	mux.Handle("DELETE /api/products/delete", middleware.Chain(http.HandlerFunc(h.DeleteProduct),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
}
