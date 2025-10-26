package categoryrepo

import models "github.com/saadahmedbd/Treestore/Models"

// CheckSlugExists checks if a slug already exists (excluding a specific ID)
func (r *categoryRepository) CheckSlugExists(slug string, excludeID uint) (bool, error) {
	var count int64
	query := r.db.Model(&models.Category{}).Where("slug = ?", slug)

	if excludeID > 0 {
		query = query.Where("id != ?", excludeID)
	}

	err := query.Count(&count).Error
	return count > 0, err
}
