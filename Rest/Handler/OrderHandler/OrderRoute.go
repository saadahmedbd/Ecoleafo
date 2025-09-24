package orderhandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *Handler) OrderRoute(mux *http.ServeMux) {
	mux.Handle("GET /getorder", middleware.Chain(http.HandlerFunc(h.GetOrder),
		middleware.Logger,
	))
	mux.Handle("GET /getorder/{orderId}", middleware.Chain(http.HandlerFunc(h.GetOrderById),
		middleware.Logger,
	))
	mux.Handle("POST /createorder", middleware.Chain(http.HandlerFunc(h.CreateOrder),
		middleware.Logger,
	))
	mux.Handle("PUT /updateorder/{orderId}", middleware.Chain(http.HandlerFunc(h.UpdateOrder),
		middleware.Logger,
	))
	mux.Handle("DELETE /deleteorder/{orderId}", middleware.Chain(http.HandlerFunc(h.DeleteOrder),
		middleware.Logger,
	))
}
