package reviewhandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *ReviewHandler) ReviewRoute(mux *http.ServeMux) {
	mux.Handle("GET /api/reviews/product",
		middleware.Chain(http.HandlerFunc(h.GetProductReviews),
			middleware.Logger,
		))
	mux.Handle("POST /api/buyer/reviews",
		middleware.Chain(http.HandlerFunc(h.CreateReview),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
		))
	mux.Handle("GET /api/buyer/reviews/my-reviews",
		middleware.Chain(http.HandlerFunc(h.GetBuyerReviews),
			middleware.Logger,
			middleware.AuthenticateJWT,
		))
	mux.Handle("GET /api/buyer/review/can-review",
		middleware.Chain(http.HandlerFunc(h.CanReview),
			middleware.Logger,
			middleware.AuthenticateJWT,
		))
	mux.Handle("DELETE /api/buyer/reviews/{id}",
		middleware.Chain(http.HandlerFunc(h.DeleteReview),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
		))
	mux.Handle("PUT /api/buyer/reviews/{id}",
		middleware.Chain(http.HandlerFunc(h.UpdateReview),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
		))
}
