package adminreviewservice

import (
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	adminreviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/AdminReviewDTO"
)

// Report review
func (s *adminReviewService) ReportReview(reviewID uint, req *adminreviewdto.ReportReviewRequest, reporterID uint, reporterType string) error {
	// Check if review exists
	_, err := s.adminreviewRepo.GetReviewByID(reviewID)
	if err != nil {
		return fmt.Errorf("review not found: %w", err)
	}

	// Create report
	report := &models.ReviewReport{
		ReviewID:     reviewID,
		ReporterID:   reporterID,
		ReporterType: reporterType,
		Reason:       req.Reason,
		Details:      req.Details,
		Status:       "pending",
	}

	return s.adminreviewRepo.CreateReviewReport(report)
}
