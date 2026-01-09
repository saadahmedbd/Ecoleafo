package commissionearningpayoutdto

type ProcessPayoutRequest struct {
	TransactionID string `json:"transaction_id" binding:"required"`
	AdminNote     string `json:"admin_note"`
}
