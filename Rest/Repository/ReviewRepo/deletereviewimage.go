package reviewrepo

import models "github.com/saadahmedbd/Treestore/Models"

// DeleteReviewImages - Delete review images
func (r *ReviewRepository) DeleteReviewImages(reviewID uint) error {
	return r.db.Where("review_id = ?", reviewID).Delete(&models.ReviewImage{}).Error
}
