package commissionearningpayoutdto

type RejectPayoutRequest struct {
	RejectionReason string `json:"rejection_reason" binding:"required"`
}
