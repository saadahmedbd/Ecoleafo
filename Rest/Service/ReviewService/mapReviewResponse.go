package reviewservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	reviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/ReviewDTO"
)

// Helper: Map review model to response DTO
func (s *ReviewService) mapReviewToResponse(review *models.Review) *reviewdto.ReviewResponse {
	response := &reviewdto.ReviewResponse{
		ID:                review.ID,
		ProductID:         review.ProductID,
		BuyerID:           review.BuyerID,
		OrderID:           review.OrderID,
		Rating:            review.Rating,
		Title:             review.Title,
		Comment:           review.Comment,
		Images:            make([]reviewdto.ReviewImageResponse, len(review.Images)),
		SellerResponse:    review.SellerResponse,
		SellerRespondedAt: review.SellerRespondedAt,
		CreatedAt:         review.CreatedAt,
		UpdatedAt:         review.UpdatedAt,
	}

	// Map images
	for i, img := range review.Images {
		response.Images[i] = reviewdto.ReviewImageResponse{
			ID:       img.ID,
			ImageURL: img.ImageURL,
			AltText:  img.AltText,
		}
	}

	// Map buyer info
	if review.Buyer.RegUser != nil {
		response.Buyer = reviewdto.BuyerInfoResponse{
			ID:             review.Buyer.ID,
			FirstName:      review.Buyer.RegUser.FirstName,
			LastName:       review.Buyer.RegUser.LastName,
			ProfilePicture: review.Buyer.ProfilePictureUrl,
		}
	}

	// Map product info
	productImage := ""
	if len(review.Product.Images) > 0 {
		productImage = review.Product.Images[0].ImageURL
	}

	response.Product = reviewdto.ProductInfoResponse{
		ID:            review.Product.ID,
		Name:          review.Product.Name,
		Slug:          review.Product.Slug,
		Image:         productImage,
		AverageRating: review.Product.AverageRating,
		ReviewCount:   review.Product.ReviewCount,
	}

	return response
}
