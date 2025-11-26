package adminreviewservice

import (
	"time"

	adminreviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/AdminReviewDTO"
)

// Get review reports
func (s *adminReviewService) GetReviewReports(status string, page, limit int) ([]adminreviewdto.ReviewReportResponse, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	reports, total, err := s.adminreviewRepo.GetReviewReports(status, page, limit)
	if err != nil {
		return nil, 0, err
	}

	var response []adminreviewdto.ReviewReportResponse
	for _, report := range reports {
		response = append(response, adminreviewdto.ReviewReportResponse{
			ID:           report.ID,
			ReviewID:     report.ReviewID,
			Review:       s.toReviewResponse(report.Review),
			ReporterType: report.ReporterType,
			Reason:       report.Reason,
			Details:      report.Details,
			Status:       report.Status,
			CreatedAt:    report.CreatedAt.Format(time.RFC3339),
		})
	}

	return response, total, nil
}
