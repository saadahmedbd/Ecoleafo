package inventory

import "time"

type StockHistoryResponse struct {
	ID        uint      `json:"id"`
	ProductID uint      `json:"product_id"`
	SellerID  uint      `json:"seller_id"`
	Type      string    `json:"type"`
	Quantity  int       `json:"quantity"`
	PrevStock int       `json:"prev_stock"`
	NewStock  int       `json:"new_stock"`
	Reason    string    `json:"reason"`
	Reference string    `json:"reference"`
	CreatedBy uint      `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}
