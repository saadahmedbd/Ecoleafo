package productrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
)

func (r *productRepository) GetProductsByCategory(categoryID uint, page, limit int) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	query := r.db.Model(&models.Product{}).
		Where("category_id = ? AND approval_status = ?", categoryID, "approved").
		Preload("Seller.RegUser").
		Preload("Category").
		Preload("Images")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&products).Error

	return products, total, err
}
