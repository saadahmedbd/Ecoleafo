package routes

import (
	"net/http"

	rolehandler "github.com/saadahmedbd/Treestore/Rest/Handler/RoleHandler"
	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func RoleRoute(mux *http.ServeMux) {
	mux.Handle("GET /getrole", middleware.Chain(http.HandlerFunc(rolehandler.GetRole),
		middleware.Logger,
	))
}
