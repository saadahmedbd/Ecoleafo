package routes

import (
	"net/http"

	orderhandler "github.com/saadahmedbd/Treestore/Handler/OrderHandler"
	middleware "github.com/saadahmedbd/Treestore/Middleware"
)

func OrderRoute(mux *http.ServeMux) {
	mux.Handle("GET /getorder", middleware.Chain(http.HandlerFunc(orderhandler.GetOrder),
		middleware.Logger,
	))
	mux.Handle("GET /getorder/{orderId}", middleware.Chain(http.HandlerFunc(orderhandler.GetOrderById),
		middleware.Logger,
	))
	mux.Handle("POST /createorder", middleware.Chain(http.HandlerFunc(orderhandler.CreateOrder),
		middleware.Logger,
	))
	mux.Handle("PUT /updateorder/{orderId}", middleware.Chain(http.HandlerFunc(orderhandler.UpdateOrder),
		middleware.Logger,
	))
	mux.Handle("DELETE /deleteorder/{orderId}", middleware.Chain(http.HandlerFunc(orderhandler.DeleteOrder),
		middleware.Logger,
	))
}
