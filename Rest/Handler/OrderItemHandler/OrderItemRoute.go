package orderitemhandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *Handler) OrderItemRoute(mux *http.ServeMux) {

	mux.Handle("GET /getorderitem", middleware.Chain(http.HandlerFunc(h.GetOrderItem),
		middleware.Logger,
	))
	mux.Handle("GET /getorderitem/{orderitemId}", middleware.Chain(http.HandlerFunc(h.GetOrderById),
		middleware.Logger,
	))
	mux.Handle("POST /createorderitem", middleware.Chain(http.HandlerFunc(h.CreateOrderItem),
		middleware.Logger,
	))
	mux.Handle("PUT /updateorderitem/{orderitemId}", middleware.Chain(http.HandlerFunc(h.UpdateOrderItem),
		middleware.Logger,
	))
	mux.Handle("DELETE /deleteorderitem/{orderitemId}", middleware.Chain(http.HandlerFunc(h.DeleteOrderItem),
		middleware.Logger,
	))

}
