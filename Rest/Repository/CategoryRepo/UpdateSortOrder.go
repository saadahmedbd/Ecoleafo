package categoryrepo

import models "github.com/saadahmedbd/Treestore/Models"

// UpdateSortOrder updates the display order of a category
func (r *categoryRepository) UpdateSortOrder(categoryID uint, sortOrder int) error {
	return r.db.Model(&models.Category{}).
		Where("id = ?", categoryID).
		Update("sort_order", sortOrder).Error
}
