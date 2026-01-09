package reviewdto

// CreateReviewRequest - DTO for creating a new review
type CreateReviewRequest struct {
	ProductID uint     `json:"product_id" binding:"required"`
	OrderID   uint     `json:"order_id" binding:"required"`
	Rating    int      `json:"rating" binding:"required,min=1,max=5"`
	Title     string   `json:"title" binding:"required,min=3,max=255"`
	Comment   string   `json:"comment" binding:"required,min=10,max=2000"`
	Images    []string `json:"images"` // Array of image URLs (uploaded separately)
}
