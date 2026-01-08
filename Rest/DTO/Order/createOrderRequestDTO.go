package order

type CreateOrderRequest struct {
	PaymentMethod    string `json:"payment_method" validate:"required,oneof=cash_on_delivery bkash card"`
	ShippingAddress  string `json:"shipping_address" validate:"required"`
	BuillingAddress  string `json:"billing_address"`
	CustomerEmail    string `json:"customer_email" validate:"required,email"`
	CustomerPhone    string `json:"customer_phone" validate:"required"`
	Notes            string `json:"notes"`
	ShippingMethodID uint   `json:"shipping_method_id"`
	CouponCode       string `json:"coupon_code"`
	IsGift           bool   `json:"is_gift"`
	GiftMessage      string `json:"gift_message"`
}
