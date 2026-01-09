package selleraccountsetting

// Toggle2FAResponse - 2FA toggle response
type Toggle2FAResponse struct {
	Enabled bool   `json:"enabled"`
	Message string `json:"message"`
}
