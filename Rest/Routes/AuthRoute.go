package routes

import (
	"net/http"

	authhandler "github.com/saadahmedbd/Treestore/Rest/Handler/AuthHandler"
	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func AuthRouth(mux *http.ServeMux) {
	mux.Handle("POST /login", middleware.Chain(http.HandlerFunc(authhandler.Login),
		middleware.Logger,
	))
	mux.Handle("POST /registation", middleware.Chain(http.HandlerFunc(authhandler.Registation),
		middleware.Logger,
	))
}
