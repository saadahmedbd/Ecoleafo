package categoryrepo

import models "github.com/saadahmedbd/Treestore/Models"

// FindByParentID retrieves all child categories of a parent
func (r *categoryRepository) FindByParentID(parentID uint) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Where("parent_id = ? AND is_active = ?", parentID, true).
		Order("sort_order ASC, name ASC").
		Find(&categories).Error
	return categories, err
}
