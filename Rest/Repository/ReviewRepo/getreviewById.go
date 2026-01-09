package reviewrepo

import models "github.com/saadahmedbd/Treestore/Models"

// GetReviewByID - Get review by ID with relations
func (r *ReviewRepository) GetReviewByID(id uint) (*models.Review, error) {
	var review models.Review
	err := r.db.Preload("Images").
		Preload("Buyer.RegUser").
		Preload("Product.Images").
		First(&review, id).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}
