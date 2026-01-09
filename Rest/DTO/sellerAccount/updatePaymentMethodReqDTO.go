package selleraccount

type UpdatePaymentMethodRequest struct {
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	BankName      string `json:"bank_name"`
	BankCode      string `json:"bank_code"`
	RoutingNumber string `json:"routing_number"`
	IsActive      bool   `json:"is_active"`
	IsDefault     bool   `json:"is_default"`
}
