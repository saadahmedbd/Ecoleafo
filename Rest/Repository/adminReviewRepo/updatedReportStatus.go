package adminreviewrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
)

// Update report status
func (r *adminReviewRepository) UpdateReportStatus(reportID uint, status string, adminID uint) error {
	now := time.Now()
	return r.db.Model(&models.ReviewReport{}).Where("id = ?", reportID).Updates(map[string]interface{}{
		"status":      status,
		"reviewed_by": adminID,
		"reviewed_at": &now,
	}).Error
}
