package adminreviewservice

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	adminreviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/AdminReviewDTO"
)

// Helper: Convert reviews to response DTOs
func (s *adminReviewService) toReviewResponses(reviews []models.Review) []adminreviewdto.ReviewResponse {
	var responses []adminreviewdto.ReviewResponse
	for _, review := range reviews {
		responses = append(responses, s.toReviewResponse(review))
	}
	return responses
}

func (s *adminReviewService) toReviewResponse(review models.Review) adminreviewdto.ReviewResponse {
	var images []string
	for _, img := range review.Images {
		images = append(images, img.ImageURL)
	}

	buyerName := ""
	buyerAvatar := ""
	if review.Buyer.RegUser != nil {
		buyerName = review.Buyer.RegUser.FirstName + " " + review.Buyer.RegUser.LastName
		buyerAvatar = review.Buyer.RegUser.Avatar
	}

	response := adminreviewdto.ReviewResponse{
		ID:                 review.ID,
		ProductID:          review.ProductID,
		ProductName:        review.Product.Name,
		BuyerID:            review.BuyerID,
		BuyerName:          buyerName,
		BuyerAvatar:        buyerAvatar,
		Rating:             review.Rating,
		Title:              review.Title,
		Comment:            review.Comment,
		Status:             review.Status,
		RejectionReason:    review.RejectionReason,
		IsVerifiedPurchase: review.IsVerifiedPurchase,
		HelpfulCount:       review.HelpfulCount,
		ReportCount:        review.ReportCount,
		IsReported:         review.IsReported,
		SellerResponse:     review.SellerResponse,
		CreatedAt:          review.CreatedAt.Format(time.RFC3339),
		Images:             images,
	}

	if review.ModeratedAt != nil {
		response.ModeratedAt = review.ModeratedAt.Format(time.RFC3339)
	}

	// Get product images
	if len(review.Product.Images) > 0 {
		response.ProductImage = review.Product.Images[0].ImageURL
	}

	// Get seller info
	response.SellerID = review.Product.SellerID
	response.SellerName = review.Product.Seller.StoreName

	return response
}
