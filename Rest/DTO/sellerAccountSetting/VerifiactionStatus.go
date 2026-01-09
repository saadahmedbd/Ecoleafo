package selleraccountsetting

// VerificationStatus - Verification status response
type VerificationStatus struct {
	Identity string `json:"identity"` // verified, pending, rejected, not-submitted
	Tax      string `json:"tax"`
	Bank     string `json:"bank"`
}
