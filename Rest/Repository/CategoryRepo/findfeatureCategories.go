package categoryrepo

import models "github.com/saadahmedbd/Treestore/Models"

// FindFeaturedCategories retrieves featured categories for homepage display
func (r *categoryRepository) FindFeaturedCategories(limit int) ([]models.Category, error) {
	var categories []models.Category
	query := r.db.Where("is_featured = ? AND is_active = ?", true, true).
		Order("sort_order ASC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&categories).Error
	return categories, err
}
