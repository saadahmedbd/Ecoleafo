package reviewservice

import (
	"fmt"

	reviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/ReviewDTO"
)

// CanReview - Check if buyer can review a product
func (s *ReviewService) CanReview(userID, productID uint) (*reviewdto.CanReviewResponse, error) {
	// Check if purchased and get buyer_id
	hasPurchased, buyerID, orderID, err := s.reviewRepo.CheckBuyerPurchased(userID, productID)
	if err != nil {
		return nil, fmt.Errorf("failed to verify purchase: %w", err)
	}

	if !hasPurchased {
		return &reviewdto.CanReviewResponse{
			CanReview:       false,
			HasPurchased:    false,
			AlreadyReviewed: false,
			Message:         "You can only review products you have purchased and received",
		}, nil
	}

	// Check if already reviewed
	alreadyReviewed, err := s.reviewRepo.CheckReviewExists(*buyerID, productID)
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

	return &reviewdto.CanReviewResponse{
		CanReview:       true,
		HasPurchased:    true,
		AlreadyReviewed: false,
		OrderID:         orderID,
		Message:         "You can review this product",
	}, nil
}
