package adminreviewservice

import adminreviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/AdminReviewDTO"

// Get review statistics
func (s *adminReviewService) GetReviewStats() (*adminreviewdto.ReviewStatsResponse, error) {
	stats, err := s.adminreviewRepo.GetReviewStats()
	if err != nil {
		return nil, err
	}

	return &adminreviewdto.ReviewStatsResponse{
		Total:              int(stats["total"].(int64)),
		Pending:            int(stats["pending"].(int64)),
		Approved:           int(stats["approved"].(int64)),
		Rejected:           int(stats["rejected"].(int64)),
		Reported:           int(stats["reported"].(int64)),
		AverageRating:      stats["average_rating"].(float64),
		RatingDistribution: stats["rating_distribution"].(map[string]int),
	}, nil
}
