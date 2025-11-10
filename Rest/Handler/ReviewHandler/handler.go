package reviewhandler

import reviewservice "github.com/saadahmedbd/Treestore/Rest/Service/ReviewService"

type ReviewHandler struct {
	reviewService *reviewservice.ReviewService
}

func NewReviewHandler(reviewService *reviewservice.ReviewService) *ReviewHandler {
	return &ReviewHandler{
		reviewService: reviewService,
	}
}
