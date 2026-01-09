package reviewservice

import reviewrepo "github.com/saadahmedbd/Treestore/Rest/Repository/ReviewRepo"

type ReviewService struct {
	reviewRepo *reviewrepo.ReviewRepository
}

func NewReviewService(reviewRepo *reviewrepo.ReviewRepository) *ReviewService {
	return &ReviewService{
		reviewRepo: reviewRepo,
	}
}
