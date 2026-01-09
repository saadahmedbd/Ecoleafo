package adminreviewservice

import adminreviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/AdminReviewDTO"

// Get pending reviews
func (s *adminReviewService) GetPendingReviews(page, limit int) ([]adminreviewdto.ReviewResponse, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	reviews, total, err := s.adminreviewRepo.GetPendingReviews(page, limit)
	if err != nil {
		return nil, 0, err
	}

	return s.toReviewResponses(reviews), total, nil
}
