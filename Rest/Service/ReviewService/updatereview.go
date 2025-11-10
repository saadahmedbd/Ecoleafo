package reviewservice

import (
	"errors"
	"fmt"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	reviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/ReviewDTO"
)

// UpdateReview - Update an existing review
func (s *ReviewService) UpdateReview(reviewID, buyerID uint, req reviewdto.UpdateReviewRequest) (*reviewdto.ReviewResponse, error) {
	// Check ownership
	isOwner, err := s.reviewRepo.CheckReviewOwnership(reviewID, buyerID)
	if err != nil {
		return nil, fmt.Errorf("failed to verify ownership: %w", err)
	}
	if !isOwner {
		return nil, errors.New("you can only update your own reviews")
	}

	// Get existing review
	review, err := s.reviewRepo.GetReviewByID(reviewID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch review: %w", err)
	}

	// Update fields
	review.Rating = req.Rating
	review.Title = req.Title
	review.Comment = req.Comment
	review.UpdatedAt = time.Now()

	if err := s.reviewRepo.UpdateReview(review); err != nil {
		return nil, fmt.Errorf("failed to update review: %w", err)
	}

	// Update images if provided
	if len(req.Images) > 0 {
		// Delete old images
		if err := s.reviewRepo.DeleteReviewImages(reviewID); err != nil {
			return nil, fmt.Errorf("failed to delete old images: %w", err)
		}

		// Create new images
		images := make([]models.ReviewImage, len(req.Images))
		for i, imgURL := range req.Images {
			images[i] = models.ReviewImage{
				ReviewID:  reviewID,
				ImageURL:  imgURL,
				AltText:   fmt.Sprintf("Review image %d", i+1),
				CreatedAt: time.Now(),
			}
		}
		if err := s.reviewRepo.CreateReviewImages(images); err != nil {
			return nil, fmt.Errorf("failed to create new images: %w", err)
		}
	}

	// Update product rating stats
	if err := s.reviewRepo.UpdateProductRatingStats(review.ProductID); err != nil {
		return nil, fmt.Errorf("failed to update product stats: %w", err)
	}

	// Fetch updated review
	updatedReview, err := s.reviewRepo.GetReviewByID(reviewID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch updated review: %w", err)
	}

	return s.mapReviewToResponse(updatedReview), nil
}
