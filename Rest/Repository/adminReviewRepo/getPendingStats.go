package adminreviewrepo

import models "github.com/saadahmedbd/Treestore/Models"

// Get product review stats
func (r *adminReviewRepository) GetProductReviewStats(productID uint) (map[string]interface{}, error) {
	var total int64
	var avgRating float64

	r.db.Model(&models.Review{}).Where("product_id = ? AND status = ?", productID, "approved").Count(&total)
	r.db.Model(&models.Review{}).Select("AVG(rating)").Where("product_id = ? AND status = ?", productID, "approved").Scan(&avgRating)

	return map[string]interface{}{
		"total_reviews":  total,
		"average_rating": avgRating,
	}, nil
}
