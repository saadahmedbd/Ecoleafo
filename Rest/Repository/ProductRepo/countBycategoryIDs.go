package productrepo

import models "github.com/saadahmedbd/Treestore/Models"

// CountByCategoryIDs counts active products across multiple categories
// Used for counting products in category + all subcategories
func (r *productRepository) CountByCategoryIDs(categoryIDs []uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Product{}).
		Where("category_id IN ? AND is_active = ?", categoryIDs, true).
		Count(&count).Error
	return count, err
}
