package authhandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *Handler) AuthRouth(mux *http.ServeMux) {
	mux.Handle("POST /api/auth/login", middleware.Chain(http.HandlerFunc(h.Login),
		middleware.Cors,
		middleware.Logger,
	))
	mux.Handle("POST /api/auth/admin/login", middleware.Chain(http.HandlerFunc(h.Login),
		middleware.Cors,
		middleware.Logger,
	))
	mux.Handle("POST /api/seller/login", middleware.Chain(http.HandlerFunc(h.Login),
		middleware.Cors,
		middleware.Logger,
	))
	mux.Handle("POST /api/auth/register", middleware.Chain(http.HandlerFunc(h.Registation),
		middleware.Cors,
		middleware.Logger,
	))
	mux.Handle("GET /api/getuser", middleware.Chain(http.HandlerFunc(h.Account),

		middleware.Logger,
	))
	mux.Handle("POST /api/auth/logout", middleware.Chain(http.HandlerFunc(h.Logout),

		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("POST /api/auth/refresh", middleware.Chain(http.HandlerFunc(h.RefreshToken),

		middleware.Logger,
		middleware.Cors,
	))
	mux.Handle("GET /api/auth/me", middleware.Chain(http.HandlerFunc(h.Me),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("POST /api/auth/logoutAll", middleware.Chain(http.HandlerFunc(h.LogoutAll),

		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	//google Oauth routes
	mux.Handle("GET /auth/google/login", middleware.Chain(http.HandlerFunc(h.GoogleLoginInit),
		middleware.Logger,
		middleware.Cors,
	))
	mux.Handle("GET /auth/google/callback", middleware.Chain(http.HandlerFunc(h.GoogleCallback),
		middleware.Logger,
		middleware.Cors,
	))
}
