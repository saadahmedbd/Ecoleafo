package order

import "time"

type OrderResponse struct {
	ID              uint                `json:"id"`
	OrderNumber     string              `json:"order_number"`
	BuyerID         uint                `json:"buyer_id"`
	Status          string              `json:"status"`
	PaymentStatus   string              `json:"payment_status"`
	PaymentMethod   string              `json:"payment_method"`
	Subtotal        float64             `json:"subtotal"`
	ShippingCost    float64             `json:"shipping_cost"`
	TaxAmount       float64             `json:"tax_amount"`
	DiscountAmount  float64             `json:"discount_amount"`
	Total           float64             `json:"total"`
	ShippingAddress string              `json:"shipping_address"`
	BillingAddress  string              `json:"billing_address"`
	CustomerEmail   string              `json:"customer_email"`
	CustomerPhone   string              `json:"customer_phone"`
	TrackingNumber  string              `json:"tracking_number"`
	ShippedAt       *time.Time          `json:"shipped_at"`
	DeliveredAt     *time.Time          `json:"delivered_at"`
	IsGift          bool                `json:"is_gift"`
	GiftMessage     string              `json:"gift_message,omitempty"`
	GiftCharge      float64             `json:"gift_charge"`
	Notes           string              `json:"notes"`
	CreatedAt       time.Time           `json:"created_at"`
	UpdatedAt       time.Time           `json:"updated_at"`
	OrderItems      []OrderItemResponse `json:"order_items,omitempty"`
}
