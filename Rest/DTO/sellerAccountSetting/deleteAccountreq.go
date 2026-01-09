package selleraccountsetting

// DeleteAccountRequest - Account deletion request
type DeleteAccountRequest struct {
	Password string `json:"password" validate:"required"`
}
