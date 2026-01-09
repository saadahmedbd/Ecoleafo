package selleraccount

type PaymentMethodInfo struct {
	ID            uint   `json:"id"`
	Type          string `json:"type"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	BankName      string `json:"bank_name"`
	IsDefault     bool   `json:"is_default"`
	IsActive      bool   `json:"is_active"`
}
