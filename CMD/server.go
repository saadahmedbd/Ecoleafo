package CMD

import (
	"fmt"
	"net/http"

	routes "github.com/saadahmedbd/Treestore/Routes"

	routehandler "github.com/saadahmedbd/Treestore/RouteHandler"
)

func Server() {
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

	// Handle all cors
	globalRoute := routehandler.GlobalHandler(mux)

	fmt.Println("3000 port server is running")
	err := http.ListenAndServe(":3000", globalRoute)
	if err != nil {
		fmt.Println("Server already used...", err)
	}
}
