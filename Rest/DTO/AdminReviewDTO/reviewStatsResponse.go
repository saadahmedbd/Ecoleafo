package adminreviewdto

// Review Statistics Response
type ReviewStatsResponse struct {
	Total              int            `json:"total"`
	Pending            int            `json:"pending"`
	Approved           int            `json:"approved"`
	Rejected           int            `json:"rejected"`
	Reported           int            `json:"reported"`
	AverageRating      float64        `json:"average_rating"`
	RatingDistribution map[string]int `json:"rating_distribution"`
}
