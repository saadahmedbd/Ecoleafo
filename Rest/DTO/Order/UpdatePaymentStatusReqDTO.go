package order

type UpdatePaymentStatusRequest struct {
	PaymentStatus string `json:"payment_status" validate:"required,oneof=pending paid failed"`
}
