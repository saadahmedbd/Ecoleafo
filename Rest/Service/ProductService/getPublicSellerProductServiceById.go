package productservice

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (s *ProductService) getProductsBySellerID(sellerID uint, filters ProductFilters, publicOnly bool) (*ProductListResponse, error) {
	// Build query
	query := s.db.Model(&models.Product{}).Where("seller_id = ?", sellerID)

	// Apply filters
	if filters.Status != "" {
		switch filters.Status {
		case "active":
			query = query.Where("is_active = ?", true)
		case "inactive":
			query = query.Where("is_active = ?", false)
		case "approved":
			query = query.Where("is_approved = ?", true)
		case "pending":
			query = query.Where("is_approved = ?", false)
		}
	}

	// For public access, always filter to active and approved products
	if publicOnly {
		query = query.Where("is_active = ? AND is_approved = ?", true, true)
	}

	if filters.Category != "" {
		if categoryID, err := strconv.ParseUint(filters.Category, 10, 32); err == nil {
			query = query.Where("category_id = ?", categoryID)
		}
	}

	if filters.Search != "" {
		searchPattern := "%" + filters.Search + "%"
		query = query.Where("name ILIKE ? OR description ILIKE ? OR scientific_name ILIKE ?",
			searchPattern, searchPattern, searchPattern)
	}

	// Get total count
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("failed to count products: %v", err)
	}

	// Apply sorting
	validSortFields := map[string]string{
		"name":       "name",
		"price":      "price",
		"created_at": "created_at",
		"updated_at": "updated_at",
		"quantity":   "quantity",
	}

	sortField, exists := validSortFields[filters.SortBy]
	if !exists {
		sortField = "created_at"
	}

	orderClause := fmt.Sprintf("%s %s", sortField, strings.ToUpper(filters.Order))
	query = query.Order(orderClause)

	// Apply pagination
	offset := (filters.Page - 1) * filters.Limit
	query = query.Offset(offset).Limit(filters.Limit)

	// Load products with relationships
	var products []models.Product
	err := query.
		Preload("Seller", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, store_name, business_email, average_rating")
		}).
		Preload("Category").
		Preload("Images", func(db *gorm.DB) *gorm.DB {
			return db.Order("sort_order ASC")
		}).
		Find(&products).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %v", err)
	}

	// Calculate pagination info
	totalPages := int(math.Ceil(float64(total) / float64(filters.Limit)))

	return &ProductListResponse{
		Products: products,
		Pagination: PaginationInfo{
			Page:       filters.Page,
			Limit:      filters.Limit,
			Total:      int(total),
			TotalPages: totalPages,
			HasNext:    filters.Page < totalPages,
			HasPrev:    filters.Page > 1,
		},
	}, nil
}
