package productrepo

import models "github.com/saadahmedbd/Treestore/Models"

// CountByCategoryID counts active products in a specific category
func (r *productRepository) CountByCategoryID(categoryID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Product{}).
		Where("category_id = ? AND is_active = ?", categoryID, true).
		Count(&count).Error
	return count, err
}
