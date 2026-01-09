package reviewrepo

import models "github.com/saadahmedbd/Treestore/Models"

// GetAverageRating - Calculate average rating for a product
func (r *ReviewRepository) GetAverageRating(productID uint) (float64, error) {
	var avgRating float64
	err := r.db.Model(&models.Review{}).
		Select("COALESCE(AVG(rating), 0)").
		Where("product_id = ?", productID).
		Scan(&avgRating).Error
	return avgRating, err
}
