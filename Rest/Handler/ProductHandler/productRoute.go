package producthandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *Handler) ProductRoute(mux *http.ServeMux) {
	mux.Handle("GET /getproduct", middleware.Chain(http.HandlerFunc(h.GetProduct),
		middleware.Logger,
	))
	mux.Handle("POST /createproduct", middleware.Chain(http.HandlerFunc(h.CreateProduct),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("GET /getproduct/{productId}", middleware.Chain(http.HandlerFunc(h.GetProductById),
		middleware.Logger,
	))
	mux.Handle("PUT /updateproduct/{productId}", middleware.Chain(http.HandlerFunc(h.UpdateProduct),
		middleware.Logger,
	))
	mux.Handle("DELETE /deleteproduct/{productId}", middleware.Chain(http.HandlerFunc(h.DeleteProduct),
		middleware.Logger,
	))
}
