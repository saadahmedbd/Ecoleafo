package commissionpayoutearningshandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *CommissionHandler) RegisterCommissionPayoutEarningRoutes(mux *http.ServeMux) {
	// Commission settings routes (admin only)

	mux.Handle("PUT /api/admin/commission/settings",
		middleware.Chain(
			http.HandlerFunc(h.UpdateCommissionSettings),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
			middleware.Requirerole([]string{"admin", "super_admin"}),
		),
	)
	mux.Handle("GET /api/admin/commission/settings",
		middleware.Chain(
			http.HandlerFunc(h.GetCommissionSettings),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
			middleware.Requirerole([]string{"admin", "super_admin"}),
		),
	)

	// Seller earnings routes (admin only)
	mux.Handle("GET /api/admin/sellers/earnings",
		middleware.Chain(
			http.HandlerFunc(h.GetSellerEarnings),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
			middleware.Requirerole([]string{"admin", "super_admin"}),
		))
	mux.Handle("GET /api/admin/sellers/earnings/all",
		middleware.Chain(http.HandlerFunc(h.GetAllSellerEarnings),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
			middleware.Requirerole([]string{"admin", "super_admin"}),
		))

	// Payout routes
	mux.Handle("POST /api/admin/payouts/request", middleware.Chain(http.HandlerFunc(h.RequestPayout),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.Requirerole([]string{"admin", "super_admin", "seller"}),
	))
	mux.Handle("GET /api/admin/payouts/pending", middleware.Chain(http.HandlerFunc(h.GetPendingPayouts),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("POST /api/admin/payouts/process", middleware.Chain(http.HandlerFunc(h.ProcessPayout),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.Requirerole([]string{"admin", "super_admin"}),
	))
	mux.Handle("POST /api/admin/payouts/reject", middleware.Chain(http.HandlerFunc(h.RejectPayout),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.Requirerole([]string{"admin", "super_admin"}),
	))
	mux.Handle("GET /api/admin/payouts/history", middleware.Chain(http.HandlerFunc(h.GetPayoutHistory),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.Requirerole([]string{"admin", "super_admin"}),
	))

	// Analytics routes (admin only)
	mux.Handle("GET /api/admin/earnings/overview", middleware.Chain(http.HandlerFunc(h.GetPlatformEarningsOverview),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.Requirerole([]string{"admin", "super_admin"}),
	))
	mux.Handle("GET /api/admin/revenue/monthly", middleware.Chain(http.HandlerFunc(h.GetMonthlyRevenue),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.Requirerole([]string{"admin", "super_admin"}),
	))
	mux.Handle("GET /api/admin/sellers/top", middleware.Chain(http.HandlerFunc(h.GetTopSellers),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.Requirerole([]string{"admin", "super_admin"}),
	))
}
