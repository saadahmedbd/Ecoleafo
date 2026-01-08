package adminreviewrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *adminReviewRepository) GetReportedReviews(page, limit int) ([]models.Review, int64, error) {
	var reviews []models.Review
	var total int64

	offset := (page - 1) * limit

	if err := r.db.Model(&models.Review{}).Where("is_reported = ?", true).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := r.db.Preload("Product.Images").
		Preload("Product.Seller").
		Preload("Buyer.RegUser").
		Preload("Images").
		Preload("Reports").
		Where("is_reported = ?", true).
		Order("report_count DESC").
		Offset(offset).
		Limit(limit).
		Find(&reviews).Error

	return reviews, total, err
}
