package order

type CreateOrderRequest struct {
	PaymentMethod         string `json:"payment_method" validate:"required,oneof=cash_on_delivery bkash card"`
	DeliveryType          string `json:"delivery_type" validate:"required,oneof=home_delivery pickup_point"`
	ShippingAddressID     *uint  `json:"shipping_address_id"`
	BillingAddressID      *uint  `json:"billing_address_id"`
	ShippingAddress       string `json:"shipping_address" validate:"required"`
	ShippingPhoneNumber   string `json:"shipping_phone_number"`
	BillingAddress        string `json:"billing_address"`
	BillingPhoneNumber    string `json:"billing_phone_number"`
	CustomerEmail         string `json:"customer_email" validate:"required,email"`
	CustomerPhone         string `json:"customer_phone" validate:"required"`
	Notes                 string `json:"notes"`
	ShippingMethodID      uint   `json:"shipping_method_id"`
	CouponCode            string `json:"coupon_code"`
	IsGift                bool   `json:"is_gift"`
	GiftMessage           string `json:"gift_message"`
}
