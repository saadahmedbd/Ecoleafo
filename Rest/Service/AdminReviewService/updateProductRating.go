package adminreviewservice

// Helper: Update product average rating
func (s *adminReviewService) updateProductRating(productID uint) {
	stats, err := s.adminreviewRepo.GetProductReviewStats(productID)
	if err != nil {
		return
	}

	s.productRepo.UpdateRating(productID, stats["average_rating"].(float64), int(stats["total_reviews"].(int64)))
}
