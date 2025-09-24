package userhandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *Handler) UserRoute(mux *http.ServeMux) {

	mux.Handle("GET /getseller", middleware.Chain(http.HandlerFunc(h.GetSeller),
		middleware.Cors,
		middleware.Logger,
	)) //Getseller route

	mux.Handle("POST /createseller", middleware.Chain(http.HandlerFunc(h.CreateSeller),
		middleware.Cors,
		middleware.Logger,
	)) // create seller route
	mux.Handle("GET /getseller/{sellerId}", middleware.Chain(http.HandlerFunc(h.GetSellerById),
		middleware.Cors,
		middleware.Logger,
	)) // get seller by id route
	mux.Handle("PUT /updateseller/{sellerId}", middleware.Chain(http.HandlerFunc(h.UpdateSeller),
		middleware.Cors,
		middleware.Logger,
	)) //update seller route
	mux.Handle("DELETE /deleteseller/{sellerId}", middleware.Chain(http.HandlerFunc(h.DeleteSeller),
		middleware.Cors,
		middleware.Logger,
	)) //update seller
}
