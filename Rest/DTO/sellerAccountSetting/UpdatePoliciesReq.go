package selleraccountsetting

type UpdatePoliciesRequest struct {
	ReturnPolicy   string `json:"return_policy" validate:"omitempty,max=2000"`
	ShippingPolicy string `json:"shipping_policy" validate:"omitempty,max=2000"`
	FAQ            string `json:"faq" validate:"omitempty,max=5000"`
}
