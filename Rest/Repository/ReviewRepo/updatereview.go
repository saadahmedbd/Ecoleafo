package reviewrepo

import models "github.com/saadahmedbd/Treestore/Models"

// UpdateReview - Update a review
func (r *ReviewRepository) UpdateReview(review *models.Review) error {
	return r.db.Save(review).Error
}
