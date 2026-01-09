package reviewdto

// ReviewImageResponse - DTO for review image
type ReviewImageResponse struct {
	ID       uint   `json:"id"`
	ImageURL string `json:"image_url"`
	AltText  string `json:"alt_text"`
}
