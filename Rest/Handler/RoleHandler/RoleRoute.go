package rolehandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *Handler) RoleRoute(mux *http.ServeMux) {
	mux.Handle("GET /getrole", middleware.Chain(http.HandlerFunc(h.GetRole),
		middleware.Logger,
	))
}
