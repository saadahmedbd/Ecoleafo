package buyerhandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *Handler) BuyerRoute(mux *http.ServeMux) {
	mux.Handle("GET /getbuyer", middleware.Chain(http.HandlerFunc(h.GetBuyer),
		middleware.Logger,
	))
	mux.Handle("POST /createbuyer", middleware.Chain(http.HandlerFunc(h.CreateBuyer),
		middleware.Logger,
	))
	mux.Handle("GET /getbuyer/{buyerId}", middleware.Chain(http.HandlerFunc(h.GetBuyerById),
		middleware.Logger,
	))
	mux.Handle("PUT /updatebuyer/{buyerId}", middleware.Chain(http.HandlerFunc(h.UpdateBuyer),
		middleware.Logger,
	))
	mux.Handle("DELETE /deletebuyer/{buyerId}", middleware.Chain(http.HandlerFunc(h.DeleteBuyer),
		middleware.Logger,
	))
}
