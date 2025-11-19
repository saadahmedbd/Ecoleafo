package commissionearningpayoutdto

type PayoutResponse struct {
	ID              uint    `json:"id"`
	SellerID        uint    `json:"seller_id"`
	SellerName      string  `json:"seller_name"`
	Amount          float64 `json:"amount"`
	PaymentMethod   string  `json:"payment_method"`
	Status          string  `json:"status"`
	TransactionID   string  `json:"transaction_id,omitempty"`
	RequestedAt     string  `json:"requested_at"`
	ProcessedAt     string  `json:"processed_at,omitempty"`
	CompletedAt     string  `json:"completed_at,omitempty"`
	RejectionReason string  `json:"rejection_reason,omitempty"`
}
