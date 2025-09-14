package routes

import (
	"net/http"

	reviewhandler "github.com/saadahmedbd/Treestore/Rest/Handler/ReviewHandler"
	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func ReviewRoute(mux *http.ServeMux) {
	mux.Handle("GET /getreview", middleware.Chain(http.HandlerFunc(reviewhandler.GetReview),
		middleware.Logger,
	))
	mux.Handle("GET /getreview/{reviewId}", middleware.Chain(http.HandlerFunc(reviewhandler.GetReviewById),
		middleware.Logger,
	))
	mux.Handle("POST /createreview", middleware.Chain(http.HandlerFunc(reviewhandler.CreateReview),
		middleware.Logger,
	))
	mux.Handle("PUT /updatereview/{reviewId}", middleware.Chain(http.HandlerFunc(reviewhandler.UpdateReview),
		middleware.Logger,
	))
	mux.Handle("DELETE /deletereview/{reviewId}", middleware.Chain(http.HandlerFunc(reviewhandler.DeleteReview),
		middleware.Logger,
	))
}
