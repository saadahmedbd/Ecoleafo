package reviewservice

import (
	"errors"
	"fmt"
	"time"

	reviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/ReviewDTO"
)

// RespondToReview - Seller responds to a review
func (s *ReviewService) RespondToReview(reviewID, userID uint, response string) (*reviewdto.ReviewResponse, error) {
	// Get seller ID from user_id
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

	// Get review
	review, err := s.reviewRepo.GetReviewByID(reviewID)
	if err != nil {
		return nil, fmt.Errorf("review not found")
	}

	// Check if user is the seller of the product
	var product struct {
		SellerID uint
	}
	err = s.reviewRepo.GetDB().Table("products").
		Select("seller_id").
		Where("id = ?", review.ProductID).
		First(&product).Error

	if err != nil {
		return nil, fmt.Errorf("product not found")
	}

	if product.SellerID != sellerID {
		return nil, errors.New("you can only respond to reviews on your own products")
	}

	// Update review with seller response
	review.SellerResponse = response
	now := time.Now()
	review.SellerRespondedAt = &now
	review.UpdatedAt = time.Now()

	if err := s.reviewRepo.UpdateReview(review); err != nil {
		return nil, fmt.Errorf("failed to add response: %w", err)
	}

	// Fetch updated review
	updatedReview, err := s.reviewRepo.GetReviewByID(reviewID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch updated review: %w", err)
	}

	return s.mapReviewToResponse(updatedReview), nil
}
