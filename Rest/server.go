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
	orderitemhandler "github.com/saadahmedbd/Treestore/Rest/Handler/OrderItemHandler"
	producthandler "github.com/saadahmedbd/Treestore/Rest/Handler/ProductHandler"
	profilehandler "github.com/saadahmedbd/Treestore/Rest/Handler/ProfileHandler"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
	routes "github.com/saadahmedbd/Treestore/Rest/Routes"
)

// crete dependency
type Server struct {
	buyerHandler     *buyerhandler.Handler
	cartitemhandler  *cartitemhandler.Handler
	categoryHandler  *categoryHandler.Handler
	orderhandler     *orderhandler.Handler
	orderitemhandler *orderitemhandler.Handler
	producthandler   *producthandler.Handler
	profilehandler   *profilehandler.Handler
}

func NewServer(
	buyerHandler *buyerhandler.Handler,
	cartitemHandler *cartitemhandler.Handler,
	categoryHandler *categoryHandler.Handler,
	orderHandler *orderhandler.Handler,
	orderitemHandler *orderitemhandler.Handler,
	productHandler *producthandler.Handler,
	profileHandler *profilehandler.Handler,
) *Server {
	return &Server{
		buyerHandler:     buyerHandler,
		cartitemhandler:  cartitemHandler,
		categoryHandler:  categoryHandler,
		orderhandler:     orderHandler,
		orderitemhandler: orderitemHandler,
		producthandler:   productHandler,
		profilehandler:   profileHandler,
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
	server.orderitemhandler.OrderItemRoute(mux)
	server.producthandler.ProductRoute(mux)
	server.profilehandler.Profile(mux)

	// handle review route
	routes.ReviewRoute(mux)

	//role route
	routes.RoleRoute(mux)
	//handle search route
	routes.Search_ProductRoute(mux)

	//handle auth route
	routes.AuthRouth(mux)

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
