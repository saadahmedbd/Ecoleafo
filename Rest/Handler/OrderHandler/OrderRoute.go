package orderhandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *OrderHandler) OrderRoute(mux *http.ServeMux) {
	//buyer route
	mux.Handle("GET /api/orders/my-orders", middleware.Chain(http.HandlerFunc(h.GetMyOrders),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("POST /api/orders/create", middleware.Chain(http.HandlerFunc(h.CreateOrder),
		middleware.AuthenticateJWT,
		middleware.Logger,
		middleware.Cors,
	))
	mux.Handle("POST /api/orders/buyer/cancel", middleware.Chain(http.HandlerFunc(h.CancelOrder),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("POST /api/orders/seller/cancel", middleware.Chain(http.HandlerFunc(h.SellerCancelOrder),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("POST /api/orders/admin/cancel", middleware.Chain(http.HandlerFunc(h.AdminCancelOrder),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	//seller route
	mux.Handle("GET /api/orders/seller-orders", middleware.Chain(http.HandlerFunc(h.GetSellerOrders),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	//common route buyer seller admin
	mux.Handle("GET /api/orders/get", middleware.Chain(http.HandlerFunc(h.GetOrderById),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("GET /api/orders/buyer/by-number", middleware.Chain(http.HandlerFunc(h.GetOrderByNumber),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("GET /api/orders/seller/by-number", middleware.Chain(http.HandlerFunc(h.GetOrderByNumberSeller),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("GET /api/orders/admin/by-number", middleware.Chain(http.HandlerFunc(h.GetOrderByNumberAdmin),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("GET /api/orders/buyer/history", middleware.Chain(http.HandlerFunc(h.GetOrderHistory),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("GET /api/orders/seller/history", middleware.Chain(http.HandlerFunc(h.GetOrderHistoryBySeller),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	//seller /admin route
	mux.Handle("POST /api/orders/update-status", middleware.Chain(http.HandlerFunc(h.UpdateOrderStatus),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("POST /api/orders/update-payment", middleware.Chain(http.HandlerFunc(h.UpdatePaymentStatus),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("POST /api/orders/update-item-status", middleware.Chain(http.HandlerFunc(h.UpdateOrderItemStatus),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	//admin route
	mux.Handle("GET /api/orders/all", middleware.Chain(http.HandlerFunc(h.GetAllOrders),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))

}
