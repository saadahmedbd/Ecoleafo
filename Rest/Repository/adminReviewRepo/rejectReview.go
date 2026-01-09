package adminreviewrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
)

// Reject review
func (r *adminReviewRepository) RejectReview(id uint, reason string, adminID uint) error {
	now := time.Now()
	return r.db.Model(&models.Review{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status":           "rejected",
		"rejection_reason": reason,
		"moderated_by":     adminID,
		"moderated_at":     &now,
	}).Error
}
