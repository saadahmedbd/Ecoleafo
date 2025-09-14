package CMD

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	middleware "github.com/saadahmedbd/Treestore/Middleware"
	routes "github.com/saadahmedbd/Treestore/Routes"
)

func Server() {
	cnf := Config.GetConfig()
	mux := http.NewServeMux()
	// router

	// user(seller) route
	routes.UserRoute(mux)

	// buyer route
	routes.BuyerRoute(mux)

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
