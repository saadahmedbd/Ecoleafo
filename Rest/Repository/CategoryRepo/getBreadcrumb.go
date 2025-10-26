package categoryrepo

import models "github.com/saadahmedbd/Treestore/Models"

// GetBreadcrumb retrieves the full category path from root to current category
// Example: Electronics > Computers > Laptops
func (r *categoryRepository) GetBreadcrumb(categoryID uint) ([]models.Category, error) {
	var breadcrumb []models.Category

	err := r.db.Raw(`
		WITH RECURSIVE category_path AS (
			SELECT *, 0 as level
			FROM categories
			WHERE id = ?
			
			UNION ALL
			
			SELECT c.*, cp.level + 1
			FROM categories c
			INNER JOIN category_path cp ON c.id = cp.parent_id
		)
		SELECT * FROM category_path
		ORDER BY level DESC
	`, categoryID).Scan(&breadcrumb).Error

	return breadcrumb, err
}
