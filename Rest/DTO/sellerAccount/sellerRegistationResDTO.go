package selleraccount

type SellerRegistrationResponse struct {
	Message       string                         `json:"message"`
	SellerID      uint                           `json:"seller_id"`
	UserID        uint                           `json:"user_id"`
	Token         string                         `json:"token"`
	NeedsSetup    bool                           `json:"needs_setup"`
	ProfileStatus *SellerProfileCompletionStatus `json:"profile_status"`
}
