package reviewservice

import (
	"fmt"
	"math"

	reviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/ReviewDTO"
)

// GetReviewsByBuyer - Get all reviews by a buyer
func (s *ReviewService) GetReviewsByBuyer(buyerID uint, page, perPage int) (*reviewdto.ReviewListResponse, error) {
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 50 {
		perPage = 10
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
		Reviews:    reviewResponses,
		Total:      total,
		Page:       page,
		PerPage:    perPage,
		TotalPages: totalPages,
	}, nil
}
