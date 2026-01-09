package adminreviewdto

// Report Review Request
type ModerateReviewRequest struct {
	Action string `json:"action" binding:"required"` // approve, reject
	Reason string `json:"reason"`                    // Required for reject
}
