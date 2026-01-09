package selleraccountsetting

// UpdateProfilePhotoResponse - Response after photo upload
type UpdateProfilePhotoResponse struct {
	PhotoURL string `json:"photo_url"`
	Message  string `json:"message"`
}
