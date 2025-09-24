package cartitemhandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *Handler) CartItemRoute(mux *http.ServeMux) {
	mux.Handle("GET /getcartitem", middleware.Chain(http.HandlerFunc(h.GetCartItem),
		middleware.Logger,
	))
	mux.Handle("GET /buyer/{buyerId}/cartcount", middleware.Chain(http.HandlerFunc(h.GetCartItemCount),
		middleware.Logger,
	))
	mux.Handle("GET /buyer/{buyerId}/cartitems", middleware.Chain(http.HandlerFunc(h.GetCartItemsByBuyer),
		middleware.Logger,
	))
}
