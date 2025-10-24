package buyerprofile

type UpdateBuyerProfileRequest struct {
	Phone          string `json:"phone"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profile_picture"`
}
