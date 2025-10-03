package buyerprofile

import "time"

type RecentOrderInfo struct {
	OrderID     uint      `json:"order_id"`
	OrderNumber string    `json:"order_number"`
	TotalAmount float64   `json:"total_amount"`
	Status      string    `json:"status"`
	ItemCount   int       `json:"item_count"`
	OrderDate   time.Time `json:"order_date"`
}
