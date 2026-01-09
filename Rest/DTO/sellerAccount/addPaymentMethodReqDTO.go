package selleraccount

type AddPaymentMethodRequest struct {
	Type          string `json:"type" validate:"required,oneof=bank_transfer bkash nagad rocket"`
	AccountName   string `json:"account_name" validate:"required"`
	AccountNumber string `json:"account_number" validate:"required"`
	BankName      string `json:"bank_name"`
	BankCode      string `json:"bank_code"`
	RoutingNumber string `json:"routing_number"`
	IsDefault     bool   `json:"is_default"`
}
