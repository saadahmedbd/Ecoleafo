package reviewservice

import (
	"fmt"

	reviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/ReviewDTO"
)

// GetSellerProductReviews - Get all reviews for seller's products
func (s *ReviewService) GetSellerProductReviews(userID, productID uint, page, perPage int) (*reviewdto.ReviewListResponse, error) {
	// Get seller ID from user_id (reg_users.id -> users.user_id)
	var seller struct {
		ID uint
	}
	err := s.reviewRepo.GetDB().Table("users").
		Select("id").
		Where("user_id = ?", userID).
		Where("deleted_at IS NULL").
		First(&seller).Error

	if err != nil {
		return nil, fmt.Errorf("seller account not found")
	}

	sellerID := seller.ID

	// Get reviews for seller's products
	reviews, total, err := s.reviewRepo.GetReviewsBySellerID(sellerID, productID, page, perPage)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch reviews: %w", err)
	}

	// Map to response
	reviewResponses := make([]reviewdto.ReviewResponse, len(reviews))
	for i, review := range reviews {
		reviewResponses[i] = *s.mapReviewToResponse(&review)
	}

	totalPages := (total + int64(perPage) - 1) / int64(perPage)

	return &reviewdto.ReviewListResponse{
		Reviews: reviewResponses,
		Pagination: reviewdto.PaginationInfo{
			Page:       page,
			PerPage:    perPage,
			Total:      total,
			TotalPages: int(totalPages),
		},
	}, nil
}
