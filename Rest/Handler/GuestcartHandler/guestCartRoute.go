package guestcarthandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *GuestcartHandler) GuestCartRoute(mux *http.ServeMux) {
	mux.Handle("GET /api/cart", middleware.Chain(http.HandlerFunc(h.GetCart),
		middleware.Cors,
		middleware.Logger,
	))
	mux.Handle("POST /api/cart/add", middleware.Chain(http.HandlerFunc(h.AddToCart),

		middleware.Cors,
		middleware.Logger,
	))
	mux.Handle("PUT /api/cart/items/{product_id}", middleware.Chain(http.HandlerFunc(h.UpdateCartItem),
		middleware.Cors,
		middleware.Logger,
	))
	mux.Handle("DELETE /api/cart/items/{product_id}", middleware.Chain(http.HandlerFunc(h.RemoveFromCart),
		middleware.Cors,
		middleware.Logger,
	))

	//protected route
	mux.Handle("GET /api/checkout/profile-status", middleware.Chain(http.HandlerFunc(h.CheckProfileStatus),
		middleware.AuthenticateJWT,
		middleware.Cors,
		middleware.Logger,
	))
	mux.Handle("POST /api/checkout/complete-profile", middleware.Chain(http.HandlerFunc(h.CompleteProfile),
		middleware.AuthenticateJWT,
		middleware.Cors,
		middleware.Logger,
	))
	mux.Handle("POST /api/checkout/address", middleware.Chain(http.HandlerFunc(h.CreateCheckoutAddress),
		middleware.AuthenticateJWT,
		middleware.Cors,
		middleware.Logger,
	))
	mux.Handle("POST /api/checkout/initiate", middleware.Chain(http.HandlerFunc(h.InitiateCheckout),
		middleware.AuthenticateJWT,
		middleware.Cors,
		middleware.Logger,
	))
	// Cart migration (called after login)
	mux.Handle("POST /api/cart/migrate", middleware.Chain(http.HandlerFunc(h.MigrateGuestCart),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))

}
