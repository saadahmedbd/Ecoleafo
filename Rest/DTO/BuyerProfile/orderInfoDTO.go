package buyerprofile

import "time"

type OrderInfo struct {
	ID          uint            `json:"id"`
	OrderNumber string          `json:"order_number"`
	TotalAmount float64         `json:"total_amount"`
	Status      string          `json:"status"`
	ItemCount   int             `json:"item_count"`
	SubTotal    float64         `json:"sub_total"`
	OrderDate   time.Time       `json:"order_date"`
	Items       []OrderItemInfo `json:"items,omitempty"`
}
