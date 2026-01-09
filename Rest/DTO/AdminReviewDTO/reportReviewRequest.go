package adminreviewdto

type ReportReviewRequest struct {
	Reason  string `json:"reason" binding:"required"` // spam, inappropriate, fake, offensive, other
	Details string `json:"details"`
}
