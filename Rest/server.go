package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	buyerhandler "github.com/saadahmedbd/Treestore/Rest/Handler/BuyerHandler"
	cartitemhandler "github.com/saadahmedbd/Treestore/Rest/Handler/CartItemHandler"
	categoryHandler "github.com/saadahmedbd/Treestore/Rest/Handler/CategoryHandler"
	orderhandler "github.com/saadahmedbd/Treestore/Rest/Handler/OrderHandler"
	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
	routes "github.com/saadahmedbd/Treestore/Rest/Routes"
)

// crete dependency
type Server struct {
	buyerHandler    *buyerhandler.Handler
	cartitemhandler *cartitemhandler.Handler
	categoryHandler *categoryHandler.Handler
	orderhandler    *orderhandler.Handler
}

func NewServer(
	buyerHandler *buyerhandler.Handler,
	cartitemHandler *cartitemhandler.Handler,
	categoryHandler *categoryHandler.Handler,
	orderHandler *orderhandler.Handler,
) *Server {
	return &Server{
		buyerHandler:    buyerHandler,
		cartitemhandler: cartitemHandler,
		categoryHandler: categoryHandler,
		orderhandler:    orderHandler,
	}

}

func (server *Server) Start(cnf Config.Config) {
	mux := http.NewServeMux()
	// router

	// user(seller) route
	routes.UserRoute(mux)

	// buyer route
	server.buyerHandler.BuyerRoute(mux)
	server.cartitemhandler.CartItemRoute(mux)
	server.categoryHandler.CategoryRoute(mux)
	server.orderhandler.OrderRoute(mux)

	//product route
	routes.ProductRoute(mux)

	//handle order item route
	routes.OrderItemRoute(mux)
	// handle review route
	routes.ReviewRoute(mux)

	//role route
	routes.RoleRoute(mux)
	//handle search route
	routes.Search_ProductRoute(mux)

	//handle auth route
	routes.AuthRouth(mux)
	//handle profile
	routes.Profile(mux)

	// Handle all cors
	handler := middleware.Cors(mux)

	address := ":" + strconv.Itoa(int(cnf.HttpPort))

	fmt.Println("server is running", address)

	err := http.ListenAndServe(address, handler)
	if err != nil {
		fmt.Println("Server already used...", err)
		os.Exit(1)
	}
}
