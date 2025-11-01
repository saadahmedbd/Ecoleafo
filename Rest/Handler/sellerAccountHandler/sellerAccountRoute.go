package selleraccounthandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *SellerRegistrationHandler) SellerAccountRoute(mux *http.ServeMux) {
	//public route
	mux.Handle("POST /api/seller/register", middleware.Chain(http.HandlerFunc(h.RegisterSeller),
		middleware.Cors,
		middleware.Logger,
	))
	//protected route
	mux.Handle("GET /api/seller/profile/status", middleware.Chain(http.HandlerFunc(h.GetProfileStatus),
		middleware.Cors,
		middleware.Logger,
		middleware.AuthenticateJWT,
	))
	mux.Handle("POST /api/seller/profile/complete", middleware.Chain(http.HandlerFunc(h.CompleteProfile),
		middleware.Cors,
		middleware.Logger,
		middleware.AuthenticateJWT,
	))
	mux.Handle("GET /api/seller-profile", middleware.Chain(http.HandlerFunc(h.GetProfile),
		middleware.Cors,
		middleware.Logger,
		middleware.AuthenticateJWT,
	))
	mux.Handle("PUT /api/seller/seller-store", middleware.Chain(http.HandlerFunc(h.UpdateStoreInfo),
		middleware.Cors,
		middleware.Logger,
		middleware.AuthenticateJWT,
	))
	//payment method route
	mux.Handle("GET /api/seller/payment-methods", middleware.Chain(http.HandlerFunc(h.GetPaymentMethods),
		middleware.Cors,
		middleware.Logger,
		middleware.AuthenticateJWT,
	))
	mux.Handle("POST /api/seller/payment-methods", middleware.Chain(http.HandlerFunc(h.AddPaymentMethod),
		middleware.Cors,
		middleware.Logger,
		middleware.AuthenticateJWT,
	))
	mux.Handle("PUT /api/seller/payment-methods/{id}", middleware.Chain(http.HandlerFunc(h.UpdatePaymentMethod),
		middleware.Cors,
		middleware.Logger,
		middleware.AuthenticateJWT,
	))
	mux.Handle("DELETE /api/seller/payment-methods/{id}", middleware.Chain(http.HandlerFunc(h.DeletePaymentMethod),
		middleware.Cors,
		middleware.Logger,
		middleware.AuthenticateJWT,
	))
	mux.Handle("PUT /api/seller/payment-methods/{id}/set-default", middleware.Chain(http.HandlerFunc(h.SetDefaultPaymentMethod),
		middleware.Cors,
		middleware.Logger,
		middleware.AuthenticateJWT,
	))
}
