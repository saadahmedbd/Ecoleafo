package reviewrepo

import models "github.com/saadahmedbd/Treestore/Models"

// GetReviewsByBuyer - Get all reviews by a buyer
func (r *ReviewRepository) GetReviewsByBuyer(buyerID uint, page, perPage int) ([]models.Review, int64, error) {
	var reviews []models.Review
	var total int64

	query := r.db.Model(&models.Review{}).Where("buyer_id = ?", buyerID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * perPage
	err := query.Offset(offset).Limit(perPage).
		Order("created_at DESC").
		Preload("Images").
		Preload("Product.Images").
		Find(&reviews).Error

	return reviews, total, err
}
