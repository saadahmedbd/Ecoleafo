package adminreviewrepo

import models "github.com/saadahmedbd/Treestore/Models"

// Get review reports
func (r *adminReviewRepository) GetReviewReports(status string, page, limit int) ([]models.ReviewReport, int64, error) {
	var reports []models.ReviewReport
	var total int64

	offset := (page - 1) * limit

	db := r.db.Model(&models.ReviewReport{})
	if status != "" {
		db = db.Where("status = ?", status)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := db.Preload("Review").
		Preload("Review.Product").
		Preload("Review.Buyer.RegUser").
		Preload("ReviewedByAdmin").
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&reports).Error

	return reports, total, err
}
