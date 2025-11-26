package adminreviewdto

// Review List Query Parameters
type ReviewListQuery struct {
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	Status     string `json:"status"` // pending, approved, rejected
	Rating     int    `json:"rating"` // 1-5
	ProductID  *uint  `json:"product_id"`
	SellerID   *uint  `json:"seller_id"`
	IsReported bool   `json:"is_reported"`
	Search     string `json:"search"`
}
