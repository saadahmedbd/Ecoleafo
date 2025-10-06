package address

type CheckoutInitRequest struct {
	ShippingAddressID uint   `json:"shipping_address_id"`
	BillingAddressID  uint   `json:"billing_address_id"`
	PaymentMethod     string `json:"payment_method"`
}
