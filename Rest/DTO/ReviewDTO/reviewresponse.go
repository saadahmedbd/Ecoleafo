package reviewdto

import "time"

// ReviewResponse - DTO for review data returned to client
type ReviewResponse struct {
	ID        uint                  `json:"id"`
	ProductID uint                  `json:"product_id"`
	BuyerID   uint                  `json:"buyer_id"`
	OrderID   *uint                 `json:"order_id"`
	Rating    int                   `json:"rating"`
	Title     string                `json:"title"`
	Comment   string                `json:"comment"`
	Images    []ReviewImageResponse `json:"images"`
	Buyer     BuyerInfoResponse     `json:"buyer"`
	Product   ProductInfoResponse   `json:"product"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}
