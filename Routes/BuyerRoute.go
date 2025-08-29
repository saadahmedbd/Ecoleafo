package routes

import (
	"net/http"

	buyerhandler "github.com/saadahmedbd/Treestore/Handler/BuyerHandler"
	middleware "github.com/saadahmedbd/Treestore/Middleware"
)

func BuyerRoute(mux *http.ServeMux) {
	mux.Handle("GET /getbuyer", middleware.Chain(http.HandlerFunc(buyerhandler.GetBuyer),
		middleware.Logger,
	))
	mux.Handle("POST /createbuyer", middleware.Chain(http.HandlerFunc(buyerhandler.CreateBuyer),
		middleware.Logger,
	))
	mux.Handle("GET /getbuyer/{buyerId}", middleware.Chain(http.HandlerFunc(buyerhandler.GetBuyerById),
		middleware.Logger,
	))
	mux.Handle("PUT /updatebuyer/{buyerId}", middleware.Chain(http.HandlerFunc(buyerhandler.UpdateBuyer),
		middleware.Logger,
	))
	mux.Handle("DELETE /deletebuyer/{buyerId}", middleware.Chain(http.HandlerFunc(buyerhandler.DeleteBuyer),
		middleware.Logger,
	))
}
