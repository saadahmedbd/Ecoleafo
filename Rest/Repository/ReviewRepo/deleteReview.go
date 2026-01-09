package reviewrepo

import models "github.com/saadahmedbd/Treestore/Models"

// DeleteReview - Soft delete a review
func (r *ReviewRepository) DeleteReview(id uint) error {
	return r.db.Delete(&models.Review{}, id).Error
}
