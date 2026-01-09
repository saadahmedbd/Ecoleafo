package sellerdashboardhandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *DashboardHandler) RegisterSellerDashboard(mux *http.ServeMux) {
	// Main statistics endpoint
	mux.Handle("GET /api/seller/statistics", middleware.Chain(
		http.HandlerFunc(h.GetStatistics),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))

	// Sales analytics
	mux.Handle("GET /api/seller/analytics/sales", middleware.Chain(
		http.HandlerFunc(h.GetSalesAnalytics),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))

	// Revenue analytics
	mux.Handle("GET /api/seller/analytics/revenue", middleware.Chain(
		http.HandlerFunc(h.GetRevenueAnalytics),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))

	// Order distribution
	mux.Handle("GET /api/seller/analytics/orders/distribution", middleware.Chain(
		http.HandlerFunc(h.GetOrderDistribution),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))

	// Performance metrics
	mux.Handle("GET /api/seller/analytics/performance", middleware.Chain(
		http.HandlerFunc(h.GetPerformanceMetrics),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))

	// Recent orders
	mux.Handle("GET /api/seller/orders/recent", middleware.Chain(
		http.HandlerFunc(h.GetRecentOrders),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))

	// Order details by ID
	mux.Handle("GET /api/seller/dashboard/orders/{id}", middleware.Chain(
		http.HandlerFunc(h.GetOrderDetails),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))

	// Update order status
	mux.Handle("PUT /api/seller/orders/{id}/status", middleware.Chain(
		http.HandlerFunc(h.UpdateOrderStatus),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))

	// Top products
	mux.Handle("GET /api/seller/products/top", middleware.Chain(
		http.HandlerFunc(h.GetTopProducts),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))

	// Low stock products
	mux.Handle("GET /api/seller/products/low-stock", middleware.Chain(
		http.HandlerFunc(h.GetLowStockProducts),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))

	// Pending actions
	mux.Handle("GET /api/seller/dashboard/pending-actions", middleware.Chain(
		http.HandlerFunc(h.GetPendingActions),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))

}
