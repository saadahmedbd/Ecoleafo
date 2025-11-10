package reviewdto

// CanReviewResponse - DTO for checking if buyer can review product
type CanReviewResponse struct {
	CanReview       bool   `json:"can_review"`
	HasPurchased    bool   `json:"has_purchased"`
	AlreadyReviewed bool   `json:"already_reviewed"`
	OrderID         *uint  `json:"order_id"`
	Message         string `json:"message"`
}
