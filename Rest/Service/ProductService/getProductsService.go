package productservice

import (
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

// PaginatedResponse - For paginated results
type PaginatedResponse struct {
	Data       []models.Product `json:"data"`
	Total      int64            `json:"total"`
	Page       int              `json:"page"`
	Limit      int              `json:"limit"`
	TotalPages int              `json:"total_pages"`
}

func (s *ProductService) GetProducts(page, limit int, filters map[string]interface{}, includeRelations bool) (*PaginatedResponse, error) {
	var product []models.Product

	var total int64
	query := s.db.Model(&models.Product{})

	//apply filter
	for key, value := range filters {
		switch key {
		case "seller_id":
			query = query.Where("seller_id = ?", value)
		case "category_id":
			query = query.Where("category_id = ?", value)
		case "is_active":
			query = query.Where("is_active = ?", value)
		case "is_approved":
			query = query.Where("is_approved = ?", value)
		case "is_featured":
			query = query.Where("is_featured = ?", value)
		case "tree_type":
			query = query.Where("tree_type = ?", value)
		case "search":
			searchTerm := fmt.Sprintf("%%%s%%", value)
			query = query.Where("name ILIKE ? OR description ILIKE ? OR sku ILIKE ?", searchTerm, searchTerm, searchTerm)
		}
	}
	query.Count(&total)
	// Always preload images
	query = query.Preload("Images", func(db *gorm.DB) *gorm.DB {
		return db.Order("sort_order ASC")
	})
	// add other preloads if requested
	if includeRelations {
		query = query.Preload("Seller").Preload("Category")
	}
	// apply pagitation
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).
		Order("created_At DESC").Find(&product).Error; err != nil {
		return nil, err
	}
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &PaginatedResponse{
		Data:       product,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}
