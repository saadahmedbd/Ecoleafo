package reviewdto

// ReviewListResponse - DTO for paginated review list
type ReviewListResponse struct {
	Reviews         []ReviewResponse `json:"reviews"`
	Total           int64            `json:"total"`
	Page            int              `json:"page"`
	PerPage         int              `json:"per_page"`
	TotalPages      int              `json:"total_pages"`
	AverageRating   float64          `json:"average_rating"`
	RatingBreakdown RatingBreakdown  `json:"rating_breakdown"`
}
