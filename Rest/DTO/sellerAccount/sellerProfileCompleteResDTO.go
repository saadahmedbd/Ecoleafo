package selleraccount

type SellerProfileCompletionStatus struct {
	IsProfileComplete bool     `json:"is_profile_complete"`
	HasBusinessInfo   bool     `json:"has_business_info"`
	HasAddress        bool     `json:"has_address"`
	HasPaymentMethod  bool     `json:"has_payment_method"`
	IsApproved        bool     `json:"is_approved"`
	IsVerified        bool     `json:"is_verified"`
	CanAddProducts    bool     `json:"can_add_products"`
	MissingFields     []string `json:"missing_fields"`
	NextStep          string   `json:"next_step"`       // "complete_profile", "add_payment", "wait_approval", "can_sell"
	ApprovalStatus    string   `json:"approval_status"` // "pending", "approved", "rejected"
	RejectionReason   string   `json:"rejection_reason,omitempty"`
}
