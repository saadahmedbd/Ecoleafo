package adminreviewrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *adminReviewRepository) GetReviewByID(ID uint) (*models.Review, error) {
	var review models.Review
	err := r.db.Preload("Product").Preload("Buyer").Preload("Buyer.RegUser").Preload("Images").Preload("ModeratedByAdmin").Preload("Reports").
		First(&review, ID).Error
	return &review, err
}
