package profilehandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *Handler) Profile(mux *http.ServeMux) {
	mux.Handle("POST /createbuyerprofile", middleware.Chain(http.HandlerFunc(h.CompleteBuyerProfile),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("POST /createsellerprofile", middleware.Chain(http.HandlerFunc(h.CompleteSeller),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
}
