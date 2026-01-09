package selleraccountsetting

// UpdateBrandingResponse - Response after logo/banner upload
type UpdateBrandingResponse struct {
	LogoURL   string `json:"logo_url,omitempty"`
	BannerURL string `json:"banner_url,omitempty"`
	Message   string `json:"message"`
}
