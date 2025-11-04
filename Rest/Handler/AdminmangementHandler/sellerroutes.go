package adminmangementhandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *SellerHandler) RegisterSeller(mux *http.ServeMux) {
	mux.Handle("GET /api/sellers", middleware.Chain(http.HandlerFunc(h.GetAllSellers),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	mux.Handle("GET /api/sellers/pending", middleware.Chain(http.HandlerFunc(h.GetPendingApprovals),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	mux.Handle("GET /api/sellers/search", middleware.Chain(http.HandlerFunc(h.SearchSellers),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	mux.Handle("GET /api/sellers/stats", middleware.Chain(http.HandlerFunc(h.GetSellerStats),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	mux.Handle("GET /api/sellers/top", middleware.Chain(http.HandlerFunc(h.GetTopSellers),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	mux.Handle("GET /api/sellers/get", middleware.Chain(http.HandlerFunc(h.GetSellerByID),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	mux.Handle("POST /api/sellers/approve", middleware.Chain(http.HandlerFunc(h.ApproveSeller),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	mux.Handle("POST /api/sellers/reject", middleware.Chain(http.HandlerFunc(h.RejectSeller),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	mux.Handle("POST /api/sellers/suspend", middleware.Chain(http.HandlerFunc(h.SuspendSeller),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	mux.Handle("POST /api/sellers/reactivate", middleware.Chain(http.HandlerFunc(h.ReactivateSeller),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))

}
