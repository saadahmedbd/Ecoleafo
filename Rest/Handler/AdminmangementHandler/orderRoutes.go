package adminmangementhandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *OrderHandler) RegisterOrderRoutes(mux *http.ServeMux) {
	mux.Handle("GET /api/orders", middleware.Chain(http.HandlerFunc(h.GetAllOrders),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	mux.Handle("GET /api/orders/search", middleware.Chain(http.HandlerFunc(h.SearchOrders),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))

	mux.Handle("GET /api/orders/stats", middleware.Chain(http.HandlerFunc(h.GetOrderStats),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))

	mux.Handle("GET /api/orders/recent", middleware.Chain(http.HandlerFunc(h.GetRecentOrders),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	//this type of roues laready have in order routes (api/orders/get)
	mux.Handle("GET /api/get/orders", middleware.Chain(http.HandlerFunc(h.GetOrderByID),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
	// //this routes alredy have into order seection
	// mux.Handle("POST /api/orders/update-status", middleware.Chain(http.HandlerFunc(h.UpdateOrderStatus),
	// 	middleware.Logger,
	// 	middleware.Cors,
	// 	middleware.AuthenticateJWT,
	// 	middleware.AdminOnlyMiddleware,
	// ))
	// // this routes alreday have in order
	// mux.Handle("POST /api/orders/cancel", middleware.Chain(http.HandlerFunc(h.CancelOrder),
	// 	middleware.Logger,
	// 	middleware.Cors,
	// 	middleware.AuthenticateJWT,
	// 	middleware.AdminOnlyMiddleware,
	// ))

}
