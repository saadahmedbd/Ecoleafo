package reviewservice

import (
	"errors"
	"fmt"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	reviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/ReviewDTO"
)

// CreateReview - Create a new product review
func (s *ReviewService) CreateReview(buyerID uint, req reviewdto.CreateReviewRequest) (*reviewdto.ReviewResponse, error) {
	// Check if buyer already reviewed this product
	exists, err := s.reviewRepo.CheckReviewExists(buyerID, req.ProductID)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing review: %w", err)
	}
	if exists {
		return nil, errors.New("you have already reviewed this product")
	}

	// Check if buyer purchased the product
	hasPurchased, orderID, err := s.reviewRepo.CheckBuyerPurchased(buyerID, req.ProductID)
	if err != nil {
		return nil, fmt.Errorf("failed to verify purchase: %w", err)
	}
	if !hasPurchased {
		return nil, errors.New("you can only review products you have purchased")
	}

	// Create review
	review := &models.Review{
		ProductID: req.ProductID,
		BuyerID:   buyerID,
		OrderID:   orderID,
		Rating:    req.Rating,
		Title:     req.Title,
		Comment:   req.Comment,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.reviewRepo.CreateReview(review); err != nil {
		return nil, fmt.Errorf("failed to create review: %w", err)
	}

	// Create review images if provided
	if len(req.Images) > 0 {
		images := make([]models.ReviewImage, len(req.Images))
		for i, imgURL := range req.Images {
			images[i] = models.ReviewImage{
				ReviewID:  review.ID,
				ImageURL:  imgURL,
				AltText:   fmt.Sprintf("Review image %d", i+1),
				CreatedAt: time.Now(),
			}
		}
		if err := s.reviewRepo.CreateReviewImages(images); err != nil {
			return nil, fmt.Errorf("failed to create review images: %w", err)
		}
	}

	// Update product rating stats
	if err := s.reviewRepo.UpdateProductRatingStats(req.ProductID); err != nil {
		return nil, fmt.Errorf("failed to update product stats: %w", err)
	}

	// Fetch complete review data
	createdReview, err := s.reviewRepo.GetReviewByID(review.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch created review: %w", err)
	}

	return s.mapReviewToResponse(createdReview), nil
}
