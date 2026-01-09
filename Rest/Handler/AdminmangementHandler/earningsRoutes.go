package adminmangementhandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *AdminManagementHandler) RegisterEarningsRoutes(mux *http.ServeMux) {
	// Seller Earnings
	mux.Handle("GET /api/admin/sellers/{id}/earnings", middleware.Chain(
		http.HandlerFunc(h.GetSellerEarningDetails),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))

	mux.Handle("GET /api/admin/sellers/earnings", middleware.Chain(
		http.HandlerFunc(h.GetAllSellerEarnings),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))

	mux.Handle("GET /api/admin/sellers/top-earning", middleware.Chain(
		http.HandlerFunc(h.GetTopEarningSellers),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))

	// Platform Earnings
	mux.Handle("GET /api/admin/platform/earnings", middleware.Chain(
		http.HandlerFunc(h.GetPlatformEarningsOverview),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))

	mux.Handle("GET /api/admin/platform/revenue/monthly", middleware.Chain(
		http.HandlerFunc(h.GetMonthlyRevenue),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
}
