package categoryrepo

import (
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
)

// FindAll retrieves categories with filtering and pagination
func (r *categoryRepository) FindAll(filter map[string]interface{}, page, limit int) ([]models.Category, int64, error) {
	var categories []models.Category
	var total int64

	// Calculate offset for pagination
	offset := (page - 1) * limit

	// Build query with filters
	query := r.db.Model(&models.Category{})

	// Apply filters
	for key, value := range filter {
		if key == "search" {
			// Search in name and description
			searchTerm := fmt.Sprintf("%%%s%%", value)
			query = query.Where("name LIKE ? OR description LIKE ?", searchTerm, searchTerm)
		} else if key == "parent_id" {
			if value == nil {
				query = query.Where("parent_id IS NULL")
			} else {
				query = query.Where("parent_id = ?", value)
			}
		} else {
			query = query.Where(fmt.Sprintf("%s = ?", key), value)
		}
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results with relationships
	err := query.Preload("Parent").
		Preload("Children").
		Order("sort_order ASC, name ASC").
		Offset(offset).
		Limit(limit).
		Find(&categories).Error

	return categories, total, err
}
