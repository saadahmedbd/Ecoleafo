package reviewdto

// UpdateReviewRequest - DTO for updating a review
type UpdateReviewRequest struct {
	Rating  int      `json:"rating" binding:"required,min=1,max=5"`
	Title   string   `json:"title" binding:"required,min=3,max=255"`
	Comment string   `json:"comment" binding:"required,min=10,max=2000"`
	Images  []string `json:"images"`
}
