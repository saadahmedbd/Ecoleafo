package orderhandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *OrderHandler) OrderRoute(mux *http.ServeMux) {
	//buyer route
	mux.Handle("GET /api/orders/my-orders", middleware.Chain(http.HandlerFunc(h.GetMyOrders),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("POST /api/orders/create", middleware.Chain(http.HandlerFunc(h.CreateOrder),
		middleware.AuthenticateJWT,
		middleware.Logger,
		middleware.Cors,
	))
	mux.Handle("POST /api/orders/cancel", middleware.Chain(http.HandlerFunc(h.CancelOrder),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	//seller route
	mux.Handle("GET /api/orders/seller-orders", middleware.Chain(http.HandlerFunc(h.GetSellerOrders),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	//common route buyer seller admin
	mux.Handle("GET /api/orders/get", middleware.Chain(http.HandlerFunc(h.GetOrderById),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("GET /api/orders/by-number", middleware.Chain(http.HandlerFunc(h.GetOrderByNumber),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("GET /api/orders/history", middleware.Chain(http.HandlerFunc(h.GetOrderHistory),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	//seller /admin route
	mux.Handle("POST /api/orders/update-status", middleware.Chain(http.HandlerFunc(h.UpdateOrderStatus),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("POST /api/orders/update-payment", middleware.Chain(http.HandlerFunc(h.UpdatePaymentStatus),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("POST /api/orders/update-item-status", middleware.Chain(http.HandlerFunc(h.UpdateOrderItemStatus),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	//admin route
	mux.Handle("GET /api/orders/all", middleware.Chain(http.HandlerFunc(h.GetAllOrders),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))

}
