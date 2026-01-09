package sellerdashboard

import "time"

type OrderDetailsResponse struct {
	ID            uint                `json:"id"`
	OrderNumber   string              `json:"order_number"`
	CustomerName  string              `json:"customer_name"`
	CustomerEmail string              `json:"customer_email"`
	CustomerPhone string              `json:"customer_phone"`
	Status        string              `json:"status"`
	PaymentStatus string              `json:"payment_status"`
	PaymentMethod string              `json:"payment_method"`
	Total         float64             `json:"total"`
	Commission    float64             `json:"commission"`
	NetEarning    float64             `json:"net_earning"`
	ShippingAddress string            `json:"shipping_address"`
	CreatedAt     time.Time           `json:"created_at"`
	Items         []OrderItemDetails  `json:"items" gorm:"-"`
}

type OrderItemDetails struct {
	ID            uint    `json:"id"`
	ProductID     uint    `json:"product_id"`
	ProductName   string  `json:"product_name"`
	ProductSKU    string  `json:"product_sku"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
	Total         float64 `json:"total"`
	Commission    float64 `json:"commission"`
	SellerEarning float64 `json:"seller_earning"`
}
