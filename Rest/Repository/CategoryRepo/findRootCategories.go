package categoryrepo

import models "github.com/saadahmedbd/Treestore/Models"

// FindRootCategories retrieves all top-level categories (no parent)
func (r *categoryRepository) FindRootCategories() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Where("parent_id IS NULL AND is_active = ?", true).
		Order("sort_order ASC, name ASC").
		Find(&categories).Error
	return categories, err
}
