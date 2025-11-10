package reviewservice

import (
	"fmt"

	reviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/ReviewDTO"
)

// CanReview - Check if buyer can review a product
func (s *ReviewService) CanReview(buyerID, productID uint) (*reviewdto.CanReviewResponse, error) {
	// Check if already reviewed
	alreadyReviewed, err := s.reviewRepo.CheckReviewExists(buyerID, productID)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing review: %w", err)
	}

	if alreadyReviewed {
		return &reviewdto.CanReviewResponse{
			CanReview:       false,
			HasPurchased:    true,
			AlreadyReviewed: true,
			Message:         "You have already reviewed this product",
		}, nil
	}

	// Check if purchased
	hasPurchased, orderID, err := s.reviewRepo.CheckBuyerPurchased(buyerID, productID)
	if err != nil {
		return nil, fmt.Errorf("failed to verify purchase: %w", err)
	}

	if !hasPurchased {
		return &reviewdto.CanReviewResponse{
			CanReview:       false,
			HasPurchased:    false,
			AlreadyReviewed: false,
			Message:         "You can only review products you have purchased",
		}, nil
	}

	return &reviewdto.CanReviewResponse{
		CanReview:       true,
		HasPurchased:    true,
		AlreadyReviewed: false,
		OrderID:         orderID,
		Message:         "You can review this product",
	}, nil
}
