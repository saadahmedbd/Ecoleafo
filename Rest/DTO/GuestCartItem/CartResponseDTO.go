package guestcartitem

type CartResponse struct {
	Items           []CartItemInfo `json:"items"`
	TotalItems      int            `json:"total_items"`
	SubTotal        float64        `json:"sub_total"`
	IsGuest         bool           `json:"is_guest"`
	CanCheckout     bool           `json:"can_checkout"`
	RequiresProfile bool           `json:"requires_profile"`
	SessionID       string         `json:"session_id,omitempty"`
}
