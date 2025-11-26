package adminreviewrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *adminReviewRepository) GetPendingReviews(page, limit int) ([]models.Review, int64, error) {
	var reviews []models.Review
	var total int64

	offset := (page - 1) * limit
	if err := r.db.Model(&models.Review{}).Where("status = ?", "pending").Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := r.db.Preload("Product").
		Preload("Buyer.RegUser").
		Preload("Images").
		Where("status = ?", "pending").
		Order("created_at ASC").
		Offset(offset).
		Limit(limit).
		Find(&reviews).Error

	return reviews, total, err

}
