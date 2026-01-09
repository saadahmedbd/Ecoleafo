package categoryrepo

import models "github.com/saadahmedbd/Treestore/Models"

// GetCategoryTree builds a complete hierarchical tree of all categories
func (r *categoryRepository) GetCategoryTree() ([]models.Category, error) {
	var categories []models.Category

	// Recursive CTE (Common Table Expression) for building tree
	// This is more efficient than multiple queries
	err := r.db.Raw(`
		WITH RECURSIVE category_tree AS (
			SELECT *, 0 as level
			FROM categories
			WHERE parent_id IS NULL AND deleted_at IS NULL AND is_active = true
			
			UNION ALL
			
			SELECT c.*, ct.level + 1
			FROM categories c
			INNER JOIN category_tree ct ON c.parent_id = ct.id
			WHERE c.deleted_at IS NULL AND c.is_active = true
		)
		SELECT * FROM category_tree
		ORDER BY level, sort_order, name
	`).Scan(&categories).Error

	return categories, err
}
