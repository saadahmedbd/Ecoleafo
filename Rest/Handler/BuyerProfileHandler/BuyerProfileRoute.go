package buyerprofilehandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *Buyerprofilehandler) BuyerRoute(mux *http.ServeMux) {
	mux.Handle("GET /api/buyer/profile", middleware.Chain(http.HandlerFunc(h.GetBuyerProfile),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("PUT /api/buyer/profile", middleware.Chain(http.HandlerFunc(h.UpdateBuyerProfile),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("PUT /api/buyer/profile/password", middleware.Chain(http.HandlerFunc(h.ChangeBuyerPassword),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("GET /api/buyer/stats", middleware.Chain(http.HandlerFunc(h.GetBuyerStats),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("GET /api/buyer/orders", middleware.Chain(http.HandlerFunc(h.GetBuyerOrderHistory),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("GET /api/buyer/addresses", middleware.Chain(http.HandlerFunc(h.GetBuyerAddress),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))

	mux.Handle("POST /api/buyer/addresses", middleware.Chain(http.HandlerFunc(h.CreateBuyerAddress),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("PUT /api/buyer/addresses/{id}", middleware.Chain(http.HandlerFunc(h.updateBuyerAddress),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("DELETE /api/buyer/addresses/{id}", middleware.Chain(http.HandlerFunc(h.DeleteAddress),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("PUT /api/buyer/addresses/{id}/set-default", middleware.Chain(http.HandlerFunc(h.SetDefaultAddress),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))

}
