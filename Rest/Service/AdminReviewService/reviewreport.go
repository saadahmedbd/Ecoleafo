package adminreviewservice

import "errors"

// Review report (admin action)
func (s *adminReviewService) ReviewReport(reportID uint, action string, adminID uint) error {
	if action != "reviewed" && action != "dismissed" {
		return errors.New("invalid action: must be 'reviewed' or 'dismissed'")
	}

	return s.adminreviewRepo.UpdateReportStatus(reportID, action, adminID)
}
