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
	))
}
