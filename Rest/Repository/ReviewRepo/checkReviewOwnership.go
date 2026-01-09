package reviewrepo

import models "github.com/saadahmedbd/Treestore/Models"

// CheckReviewOwnership - Check if review belongs to buyer
func (r *ReviewRepository) CheckReviewOwnership(reviewID, buyerID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.Review{}).
		Where("id = ? AND buyer_id = ?", reviewID, buyerID).
		Count(&count).Error
	return count > 0, err
}
