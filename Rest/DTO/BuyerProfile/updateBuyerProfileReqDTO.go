package buyerprofile

type UpdateBuyerProfileRequest struct {
	Phone          string `json:"phone"`
	DefaultAddress string `json:"default_address"`
}
