package reviewrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	reviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/ReviewDTO"
)

// GetReviewsByProduct - Get all reviews for a product with pagination
func (r *ReviewRepository) GetReviewsByProduct(filter reviewdto.ReviewFilterRequest) ([]models.Review, int64, error) {
	var reviews []models.Review
	var total int64

	query := r.db.Model(&models.Review{}).Where("product_id = ?", filter.ProductID)

	// Apply rating filter
	if filter.Rating > 0 {
		query = query.Where("rating = ?", filter.Rating)
	}

	// Count total
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply sorting
	switch filter.SortBy {
	case "oldest":
		query = query.Order("created_at ASC")
	case "highest_rated":
		query = query.Order("rating DESC, created_at DESC")
	case "lowest_rated":
		query = query.Order("rating ASC, created_at DESC")
	default: // newest
		query = query.Order("created_at DESC")
	}

	// Apply pagination
	offset := (filter.Page - 1) * filter.PerPage
	err := query.Offset(offset).Limit(filter.PerPage).
		Preload("Images").
		Preload("Buyer.RegUser").
		Preload("Product").
		Find(&reviews).Error

	return reviews, total, err
}
