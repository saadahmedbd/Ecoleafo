package reviewrepo

import models "github.com/saadahmedbd/Treestore/Models"

// CheckReviewExists - Check if buyer already reviewed the product
func (r *ReviewRepository) CheckReviewExists(buyerID, productID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.Review{}).
		Where("buyer_id = ? AND product_id = ?", buyerID, productID).
		Count(&count).Error
	return count > 0, err
}
