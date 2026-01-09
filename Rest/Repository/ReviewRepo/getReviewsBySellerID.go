package reviewrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
)

// GetReviewsBySellerID - Get all reviews for seller's products
func (r *ReviewRepository) GetReviewsBySellerID(sellerID, productID uint, page, perPage int) ([]models.Review, int64, error) {
	var reviews []models.Review
	var total int64

	offset := (page - 1) * perPage

	query := r.db.Model(&models.Review{}).
		Joins("JOIN products ON products.id = reviews.product_id").
		Where("products.seller_id = ?", sellerID).
		Where("reviews.deleted_at IS NULL")

	// Filter by specific product if provided
	if productID > 0 {
		query = query.Where("reviews.product_id = ?", productID)
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get reviews with preloaded data
	err := query.
		Preload("Buyer.RegUser").
		Preload("Product.Images").
		Preload("Images").
		Order("reviews.created_at DESC").
		Offset(offset).
		Limit(perPage).
		Find(&reviews).Error

	return reviews, total, err
}
