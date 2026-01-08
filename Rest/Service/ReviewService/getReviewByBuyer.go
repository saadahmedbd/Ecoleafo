package reviewservice

import (
	"fmt"
	"math"

	reviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/ReviewDTO"
)

// GetReviewsByBuyer - Get all reviews by a buyer
func (s *ReviewService) GetReviewsByBuyer(userID uint, page, perPage int) (*reviewdto.ReviewListResponse, error) {
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 50 {
		perPage = 10
	}

	// Get buyer_id from user_id
	buyerID, err := s.reviewRepo.GetBuyerIDByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("buyer account not found: %w", err)
	}

	// Debug: Check if buyerID is valid
	if buyerID == 0 {
		return nil, fmt.Errorf("invalid buyer ID: %d", buyerID)
	}

	reviews, total, err := s.reviewRepo.GetReviewsByBuyer(buyerID, page, perPage)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch buyer reviews: %w", err)
	}

	reviewResponses := make([]reviewdto.ReviewResponse, len(reviews))
	for i, review := range reviews {
		reviewResponses[i] = *s.mapReviewToResponse(&review)
	}

	totalPages := int(math.Ceil(float64(total) / float64(perPage)))

	return &reviewdto.ReviewListResponse{
		Reviews: reviewResponses,
		Pagination: reviewdto.PaginationInfo{
			Page:       page,
			PerPage:    perPage,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}
