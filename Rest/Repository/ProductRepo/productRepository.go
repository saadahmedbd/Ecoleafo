package productrepo

// ============================================================================
// PRODUCT REPOSITORY (Minimal - Only for Category System)
// ============================================================================

import (
	"gorm.io/gorm"
)

// ProductRepository interface with only methods needed for category system
type ProductRepository interface {
	// Method used by category service to count products
	CountByCategoryID(categoryID uint) (int64, error)

	// Method to count products in multiple categories at once
	CountByCategoryIDs(categoryIDs []uint) (int64, error)
}

type productRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new product repository instance
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}
