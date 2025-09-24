package reviewhandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *Handler) ReviewRoute(mux *http.ServeMux) {
	mux.Handle("GET /getreview", middleware.Chain(http.HandlerFunc(h.GetReview),
		middleware.Logger,
	))
	mux.Handle("GET /getreview/{reviewId}", middleware.Chain(http.HandlerFunc(h.GetReviewById),
		middleware.Logger,
	))
	mux.Handle("POST /createreview", middleware.Chain(http.HandlerFunc(h.CreateReview),
		middleware.Logger,
	))
	mux.Handle("PUT /updatereview/{reviewId}", middleware.Chain(http.HandlerFunc(h.UpdateReview),
		middleware.Logger,
	))
	mux.Handle("DELETE /deletereview/{reviewId}", middleware.Chain(http.HandlerFunc(h.DeleteReview),
		middleware.Logger,
	))
}
