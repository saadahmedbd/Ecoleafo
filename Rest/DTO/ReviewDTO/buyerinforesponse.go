package reviewdto

// BuyerInfoResponse - DTO for buyer information in review
type BuyerInfoResponse struct {
	ID             uint   `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	ProfilePicture string `json:"profile_picture"`
}
