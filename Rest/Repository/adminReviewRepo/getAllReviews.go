package adminreviewrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	adminreviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/AdminReviewDTO"
)

// Get all reviews with filters
func (r *adminReviewRepository) GetAllReviews(query adminreviewdto.ReviewListQuery) ([]models.Review, int64, error) {
	var reviews []models.Review
	var total int64

	db := r.db.Model(&models.Review{})

	// Apply filters
	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}
	if query.Rating > 0 {
		db = db.Where("rating = ?", query.Rating)
	}
	if query.ProductID != nil {
		db = db.Where("product_id = ?", *query.ProductID)
	}
	if query.IsReported {
		db = db.Where("is_reported = ?", true)
	}
	if query.Search != "" {
		db = db.Where("title ILIKE ? OR comment ILIKE ?", "%"+query.Search+"%", "%"+query.Search+"%")
	}

	// Get total count
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	offset := (query.Page - 1) * query.Limit
	err := db.Preload("Product").
		Preload("Buyer").
		Preload("Buyer.RegUser").
		Preload("Images").
		Preload("ModeratedByAdmin").
		Order("created_at DESC").
		Offset(offset).
		Limit(query.Limit).
		Find(&reviews).Error

	return reviews, total, err
}
