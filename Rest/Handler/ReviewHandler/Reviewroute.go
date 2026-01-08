package reviewhandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *ReviewHandler) ReviewRoute(mux *http.ServeMux) {
	// Public routes
	mux.Handle("GET /api/reviews/product",
		middleware.Chain(http.HandlerFunc(h.GetProductReviews),
			middleware.Logger,
		))

	// Buyer routes
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
			middleware.Cors,
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

	// Seller routes
	mux.Handle("GET /api/seller/reviews/my-products",
		middleware.Chain(http.HandlerFunc(h.GetSellerProductReviews),
			middleware.Logger,
			middleware.AuthenticateJWT,
		))
	mux.Handle("POST /api/seller/reviews/{id}/respond",
		middleware.Chain(http.HandlerFunc(h.RespondToReview),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
		))
}
