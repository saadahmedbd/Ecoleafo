package adminreviewrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

// Create review report
func (r *adminReviewRepository) CreateReviewReport(report *models.ReviewReport) error {
	if err := r.db.Create(report).Error; err != nil {
		return err
	}

	// Update review's is_reported flag and increment report count
	return r.db.Model(&models.Review{}).Where("id = ?", report.ReviewID).Updates(map[string]interface{}{
		"is_reported":  true,
		"report_count": gorm.Expr("report_count + 1"),
	}).Error
}
