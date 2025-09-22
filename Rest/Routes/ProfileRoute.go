package routes

import (
	"net/http"

	profilehandler "github.com/saadahmedbd/Treestore/Rest/Handler/ProfileHandler"
	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func Profile(mux *http.ServeMux) {
	mux.Handle("POST /createbuyerprofile", middleware.Chain(http.HandlerFunc(profilehandler.CompleteBuyerProfile),
		middleware.Logger,
	))
}
