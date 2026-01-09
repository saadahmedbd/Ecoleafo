package inventoryhandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *Inventoryhandler) RegisterInventoryRoute(mux *http.ServeMux) {
	mux.Handle("GET /api/seller/inventory", middleware.Chain(
		http.HandlerFunc(h.GetInventory),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("GET /api/seller/inventory/stats", middleware.Chain(
		http.HandlerFunc(h.GetInventoryStats),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("PUT /api/seller/inventory/{productId}/stock", middleware.Chain(
		http.HandlerFunc(h.UpdateStock),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("POST /api/seller/inventory/bulk-update", middleware.Chain(
		http.HandlerFunc(h.BulkUpdateStock),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("PUT /api/seller/inventory/{productId}/thresold", middleware.Chain(
		http.HandlerFunc(h.UpdateThreshold),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("GET /api/seller/inventory/{productId}/history", middleware.Chain(
		http.HandlerFunc(h.GetStockHistory),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("GET /api/seller/inventory/export", middleware.Chain(
		http.HandlerFunc(h.ExportInventory),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
}
