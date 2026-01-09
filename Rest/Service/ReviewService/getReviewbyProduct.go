package reviewservice

import (
	"fmt"
	"math"

	reviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/ReviewDTO"
)

// GetReviewsByProduct - Get all reviews for a product
func (s *ReviewService) GetReviewsByProduct(filter reviewdto.ReviewFilterRequest) (*reviewdto.ReviewListResponse, error) {
	// Set default pagination
	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.PerPage < 1 || filter.PerPage > 50 {
		filter.PerPage = 10
	}

	reviews, total, err := s.reviewRepo.GetReviewsByProduct(filter)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch reviews: %w", err)
	}

	// Map to response
	reviewResponses := make([]reviewdto.ReviewResponse, len(reviews))
	for i, review := range reviews {
		reviewResponses[i] = *s.mapReviewToResponse(&review)
	}

	totalPages := int(math.Ceil(float64(total) / float64(filter.PerPage)))

	return &reviewdto.ReviewListResponse{
		Reviews: reviewResponses,
		Pagination: reviewdto.PaginationInfo{
			Page:       filter.Page,
			PerPage:    filter.PerPage,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}
