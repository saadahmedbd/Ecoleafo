package authhandler

import (
	"net/http"


	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *Handler)AuthRouth(mux *http.ServeMux) {
	mux.Handle("POST /login", middleware.Chain(http.HandlerFunc(h.Login),
		middleware.Logger,
	))
	mux.Handle("POST /registation", middleware.Chain(http.HandlerFunc(h.Registation),
		middleware.Cors,
		middleware.Logger,
	))
	mux.Handle("GET /getuser", middleware.Chain(http.HandlerFunc(h.Account),

		middleware.Logger,
	))
}
