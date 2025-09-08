package CMD

import (
	"fmt"
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Middleware"
	routes "github.com/saadahmedbd/Treestore/Routes"
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
	//handle search route
	routes.Search_ProductRoute(mux)

	// Handle all cors
	handler := middleware.Cors(mux)

	fmt.Println("3000 port server is running")
	err := http.ListenAndServe(":3000", handler)
	if err != nil {
		fmt.Println("Server already used...", err)
	}
}
