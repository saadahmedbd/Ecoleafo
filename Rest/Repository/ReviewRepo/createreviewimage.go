package reviewrepo

import models "github.com/saadahmedbd/Treestore/Models"

// CreateReviewImages - Create review images
func (r *ReviewRepository) CreateReviewImages(images []models.ReviewImage) error {
	if len(images) == 0 {
		return nil
	}
	return r.db.Create(&images).Error
}
