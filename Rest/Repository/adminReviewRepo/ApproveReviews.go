package adminreviewrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
)

// Approve review
func (r *adminReviewRepository) ApproveReview(id uint, adminID uint) error {
	now := time.Now()
	return r.db.Model(&models.Review{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status":       "approved",
		"moderated_by": adminID,
		"moderated_at": &now,
	}).Error
}
