package reviewrepo

import models "github.com/saadahmedbd/Treestore/Models"

// CreateReview - Create a new review
func (r *ReviewRepository) CreateReview(review *models.Review) error {
	return r.db.Create(review).Error
}
