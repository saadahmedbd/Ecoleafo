package buyerprofile

type UpdateProfilePhotoResponse struct {
	PhotoURL string `json:"photo_url"`
	PublicID string `json:"public_id"`
	Message  string `json:"message"`
}
