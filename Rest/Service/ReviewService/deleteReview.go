package reviewservice

import (
	"errors"
	"fmt"
)

// DeleteReview - Delete a review
func (s *ReviewService) DeleteReview(reviewID, buyerID uint) error {
	// Check ownership
	isOwner, err := s.reviewRepo.CheckReviewOwnership(reviewID, buyerID)
	if err != nil {
		return fmt.Errorf("failed to verify ownership: %w", err)
	}
	if !isOwner {
		return errors.New("you can only delete your own reviews")
	}

	// Get review to update product stats later
	review, err := s.reviewRepo.GetReviewByID(reviewID)
	if err != nil {
		return fmt.Errorf("failed to fetch review: %w", err)
	}

	// Delete review images
	if err := s.reviewRepo.DeleteReviewImages(reviewID); err != nil {
		return fmt.Errorf("failed to delete review images: %w", err)
	}

	// Delete review
	if err := s.reviewRepo.DeleteReview(reviewID); err != nil {
		return fmt.Errorf("failed to delete review: %w", err)
	}

	// Update product rating stats
	if err := s.reviewRepo.UpdateProductRatingStats(review.ProductID); err != nil {
		return fmt.Errorf("failed to update product stats: %w", err)
	}

	return nil
}
