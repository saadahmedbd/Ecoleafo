package reviewdto

// ReviewFilterRequest - DTO for filtering reviews
type ReviewFilterRequest struct {
	ProductID uint   `form:"product_id"`
	BuyerID   uint   `form:"buyer_id"`
	Rating    int    `form:"rating"`
	Page      int    `form:"page"`
	PerPage   int    `form:"per_page"`
	SortBy    string `form:"sort_by"` // newest, oldest, highest_rated, lowest_rated
}
