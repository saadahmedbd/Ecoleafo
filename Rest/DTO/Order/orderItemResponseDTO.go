package order

import "time"

type OrderItemResponse struct {
	ID            uint       `json:"id"`
	ProductID     uint       `json:"product_id"`
	ProductName   string     `json:"product_name"`
	ProductSKU    string     `json:"product_sku"`
	SellerID      uint       `json:"seller_id"`
	SellerName    string     `json:"seller_name"`
	Quantity      int        `json:"quantity"`
	Image         string     `json:"image"`
	Price         float64    `json:"price"`
	Total         float64    `json:"total"`
	Commission    float64    `json:"commission"`
	SellerEarning float64    `json:"seller_earning"`
	Status        string     `json:"status"`
	ShippedAt     *time.Time `json:"shipped_at"`
	DeliveredAt   *time.Time `json:"delivered_at"`
}
