package reviewrepo

import models "github.com/saadahmedbd/Treestore/Models"

// UpdateProductRatingStats - Update product's average rating and review count
func (r *ReviewRepository) UpdateProductRatingStats(productID uint) error {
	avgRating, err := r.GetAverageRating(productID)
	if err != nil {
		return err
	}

	reviewCount, err := r.GetReviewCount(productID)
	if err != nil {
		return err
	}

	return r.db.Model(&models.Product{}).
		Where("id = ?", productID).
		Updates(map[string]interface{}{
			"average_rating": avgRating,
			"review_count":   reviewCount,
		}).Error
}
