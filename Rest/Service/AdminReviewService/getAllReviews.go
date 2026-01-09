package adminreviewservice

import adminreviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/AdminReviewDTO"

// Get all reviews
func (s *adminReviewService) GetAllReviews(query adminreviewdto.ReviewListQuery) ([]adminreviewdto.ReviewResponse, int64, error) {
	if query.Page < 1 {
		query.Page = 1
	}
	if query.Limit < 1 || query.Limit > 100 {
		query.Limit = 20
	}

	reviews, total, err := s.adminreviewRepo.GetAllReviews(query)
	if err != nil {
		return nil, 0, err
	}

	return s.toReviewResponses(reviews), total, nil
}
