package reviewrepo

import models "github.com/saadahmedbd/Treestore/Models"

// GetReviewCount - Get total review count for a product
func (r *ReviewRepository) GetReviewCount(productID uint) (int, error) {
	var count int64
	err := r.db.Model(&models.Review{}).
		Where("product_id = ?", productID).
		Count(&count).Error
	return int(count), err
}
