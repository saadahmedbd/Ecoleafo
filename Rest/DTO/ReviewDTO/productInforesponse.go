package reviewdto

// ProductInfoResponse - DTO for product information in review
type ProductInfoResponse struct {
	ID            uint    `json:"id"`
	Name          string  `json:"name"`
	Slug          string  `json:"slug"`
	Image         string  `json:"image"` // First product image
	AverageRating float64 `json:"average_rating"`
	ReviewCount   int     `json:"review_count"`
}
