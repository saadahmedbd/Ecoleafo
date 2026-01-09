package productrepo

import models "github.com/saadahmedbd/Treestore/Models"

// UpdateRating updates product's average rating and review count
func (r *productRepository) UpdateRating(productID uint, averageRating float64, reviewCount int) error {
	return r.db.Model(&models.Product{}).
		Where("id = ?", productID).
		Updates(map[string]interface{}{
			"average_rating": averageRating,
			"review_count":   reviewCount,
		}).Error
}
