package adminreviewrepo

import models "github.com/saadahmedbd/Treestore/Models"

// Get reports by review ID
func (r *adminReviewRepository) GetReportsByReviewID(reviewID uint) ([]models.ReviewReport, error) {
	var reports []models.ReviewReport
	err := r.db.Where("review_id = ?", reviewID).
		Order("created_at DESC").
		Find(&reports).Error
	return reports, err
}
