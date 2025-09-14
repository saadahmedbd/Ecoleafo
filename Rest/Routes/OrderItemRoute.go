package routes

import (
	"net/http"

	orderitemhandler "github.com/saadahmedbd/Treestore/Rest/Handler/OrderItemHandler"
	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func OrderItemRoute(mux *http.ServeMux) {

	mux.Handle("GET /getorderitem", middleware.Chain(http.HandlerFunc(orderitemhandler.GetOrderItem),
		middleware.Logger,
	))
	mux.Handle("GET /getorderitem/{orderitemId}", middleware.Chain(http.HandlerFunc(orderitemhandler.GetOrderById),
		middleware.Logger,
	))
	mux.Handle("POST /createorderitem", middleware.Chain(http.HandlerFunc(orderitemhandler.CreateOrderItem),
		middleware.Logger,
	))
	mux.Handle("PUT /updateorderitem/{orderitemId}", middleware.Chain(http.HandlerFunc(orderitemhandler.UpdateOrderItem),
		middleware.Logger,
	))
	mux.Handle("DELETE /deleteorderitem/{orderitemId}", middleware.Chain(http.HandlerFunc(orderitemhandler.DeleteOrderItem),
		middleware.Logger,
	))

}
