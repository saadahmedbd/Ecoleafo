package adminreviewdto

// Review Response DTO
type ReviewResponse struct {
	ID                 uint     `json:"id"`
	ProductID          uint     `json:"product_id"`
	ProductName        string   `json:"product_name"`
	ProductImage       string   `json:"product_image"`
	BuyerID            uint     `json:"buyer_id"`
	BuyerName          string   `json:"buyer_name"`
	BuyerAvatar        string   `json:"buyer_avatar"` //buyer profile picture
	SellerID           uint     `json:"seller_id"`
	SellerName         string   `json:"seller_name"`
	Rating             int      `json:"rating"`
	Title              string   `json:"title"`
	Comment            string   `json:"comment"`
	Status             string   `json:"status"`
	RejectionReason    string   `json:"rejection_reason,omitempty"`
	IsVerifiedPurchase bool     `json:"is_verified_purchase"`
	HelpfulCount       int      `json:"helpful_count"`
	ReportCount        int      `json:"report_count"`
	IsReported         bool     `json:"is_reported"`
	SellerResponse     string   `json:"seller_response,omitempty"`
	CreatedAt          string   `json:"created_at"`
	ModeratedAt        string   `json:"moderated_at,omitempty"`
	Images             []string `json:"images"`
}
