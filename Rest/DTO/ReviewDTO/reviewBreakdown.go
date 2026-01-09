package reviewdto

// RatingBreakdown - DTO for rating statistics
type RatingBreakdown struct {
	FiveStar  int `json:"five_star"`
	FourStar  int `json:"four_star"`
	ThreeStar int `json:"three_star"`
	TwoStar   int `json:"two_star"`
	OneStar   int `json:"one_star"`
}
