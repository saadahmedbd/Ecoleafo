package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	buyerhandler "github.com/saadahmedbd/Treestore/Rest/Handler/BuyerHandler"
	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
	routes "github.com/saadahmedbd/Treestore/Rest/Routes"
)

// crete dependency
type Server struct {
	buyerHandler *buyerhandler.Handler
}

func NewServer(
	buyerHandler *buyerhandler.Handler,
) *Server {
	return &Server{
		buyerHandler: buyerHandler,
	}

}

func (server *Server) Start(cnf Config.Config) {
	mux := http.NewServeMux()
	// router

	// user(seller) route
	routes.UserRoute(mux)

	// buyer route
	server.buyerHandler.BuyerRoute(mux)

	//product route
	routes.ProductRoute(mux)

	//category route
	routes.CategoryRoute(mux)

	// handle cartitem route
	routes.CartItemRoute(mux)

	// handle order route
	routes.OrderRoute(mux)
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
