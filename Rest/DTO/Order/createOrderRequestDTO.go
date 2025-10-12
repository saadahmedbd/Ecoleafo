package order

type CreateOrderRequest struct {
	PaymentMethod   string  `json:"payment_method" validate:"required,oneof=cash_on_delivery bkash card"`
	ShippingAddress string  `json:"shipping_address" validate:"required"`
	BuillingAddress string  `json:"billing_address"`
	CustomerEmail   string  `json:"customer_email" validate:"required,email"`
	CustomerPhone   string  `json:"customer_phone" validate:"required"`
	ShippingCost    float64 `json:"shipping_cost"`
	DiscountAmount  float64 `json:"discount_amount"`
	Notes           string  `json:"notes"`
}
