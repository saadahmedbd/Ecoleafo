package adminmangementhandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *BuyerHandler) RegisterBuyer(mux *http.ServeMux) {
	mux.Handle("GET /api/buyers", middleware.Chain(http.HandlerFunc(h.GetAllBuyers),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	mux.Handle("GET /api/buyers/search", middleware.Chain(http.HandlerFunc(h.SearchBuyers),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	mux.Handle("GET /api/buyers/stats", middleware.Chain(http.HandlerFunc(h.GetBuyerStats),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	mux.Handle("GET /api/buyers/get", middleware.Chain(http.HandlerFunc(h.GetBuyerByID),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	mux.Handle("POST /api/buyers/activate", middleware.Chain(http.HandlerFunc(h.ActivateBuyer),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	mux.Handle("POST /api/buyers/deactivate", middleware.Chain(http.HandlerFunc(h.DeactivateBuyer),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	mux.Handle("POST /api/buyers/suspend", middleware.Chain(http.HandlerFunc(h.SuspendBuyer),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
}
