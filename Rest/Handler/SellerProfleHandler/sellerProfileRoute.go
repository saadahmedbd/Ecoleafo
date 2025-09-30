package sellerproflehandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *SellerProfileHandler) SellerRoute(mux *http.ServeMux) {
	mux.Handle("GET /api/seller/profile", middleware.Chain(http.HandlerFunc(h.GetSellerProfile),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("PUT /api/seller/profile", middleware.Chain(http.HandlerFunc(h.UpdateSellerProfile),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("PUT /api/seller/profile/image", middleware.Chain(http.HandlerFunc(h.UpdateStoreImage),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("PUT /api/seller/profile/password", middleware.Chain(http.HandlerFunc(h.ChangeSellerPassword),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("GET /api/seller/state", middleware.Chain(http.HandlerFunc(h.GetSellerState),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("GET /api/seller/{id}/profile", middleware.Chain(http.HandlerFunc(h.GetPublicSellerProfile),
		middleware.Logger,
	))
}
