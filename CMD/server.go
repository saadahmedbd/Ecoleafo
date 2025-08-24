package CMD

import (
	"fmt"
	"net/http"

	rolehandler "github.com/saadahmedbd/Treestore/Handler/RoleHandler"
	userhandler "github.com/saadahmedbd/Treestore/Handler/UserHandler"
	routehandler "github.com/saadahmedbd/Treestore/RouteHandler"
)

func Server() {
	mux := http.NewServeMux()

	// router
	mux.HandleFunc("/home", userhandler.Home)
	// user(seller) route
	mux.Handle("GET /getseller", http.HandlerFunc(userhandler.GetSeller))        //Getseller route
	mux.Handle("POST /createseller", http.HandlerFunc(userhandler.CreateSeller)) // create seller route

	//role route
	mux.Handle("GET /getrole", http.HandlerFunc(rolehandler.GetRole))

	// Handle all cors
	globalRoute := routehandler.GlobalHandler(mux)

	fmt.Println("3000 port server is running")
	err := http.ListenAndServe(":3000", globalRoute)
	if err != nil {
		fmt.Println("Server already used...", err)
	}
}
