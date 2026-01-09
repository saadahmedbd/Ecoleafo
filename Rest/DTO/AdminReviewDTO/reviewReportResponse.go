package adminreviewdto

// Review Report Response
type ReviewReportResponse struct {
	ID           uint           `json:"id"`
	ReviewID     uint           `json:"review_id"`
	Review       ReviewResponse `json:"review"`
	ReporterType string         `json:"reporter_type"`
	Reason       string         `json:"reason"`
	Details      string         `json:"details"`
	Status       string         `json:"status"`
	CreatedAt    string         `json:"created_at"`
}
