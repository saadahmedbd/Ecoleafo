package adminreviewrepo

import models "github.com/saadahmedbd/Treestore/Models"

// Delete review (soft delete)
func (r *adminReviewRepository) DeleteReview(id uint) error {
	return r.db.Delete(&models.Review{}, id).Error
}
