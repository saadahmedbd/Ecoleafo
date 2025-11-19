package commissionearningpayoutdto

type CreatePayoutRequest struct {
	SellerID      uint    `json:"seller_id" binding:"required"`
	Amount        float64 `json:"amount" binding:"required,min=1"`
	PaymentMethod string  `json:"payment_method" binding:"required"`
	BankName      string  `json:"bank_name"`
	AccountNumber string  `json:"account_number"`
	AccountName   string  `json:"account_name"`
	RequestNote   string  `json:"request_note"`
}
