package routes

import (
	"net/http"

	cartItemHandler "github.com/saadahmedbd/Treestore/Rest/Handler/CartItemHandler"
	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func CartItemRoute(mux *http.ServeMux) {
	mux.Handle("GET /getcartitem", middleware.Chain(http.HandlerFunc(cartItemHandler.GetCartItem),
		middleware.Logger,
	))
	mux.Handle("GET /buyer/{buyerId}/cartcount", middleware.Chain(http.HandlerFunc(cartItemHandler.GetCartItemCount),
		middleware.Logger,
	))
	mux.Handle("GET /buyer/{buyerId}/cartitems", middleware.Chain(http.HandlerFunc(cartItemHandler.GetCartItemsByBuyer),
		middleware.Logger,
	))
}
