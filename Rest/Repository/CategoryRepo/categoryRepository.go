package categoryrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	// Basic CRUD operations
	Create(category *models.Category) error
	FindByID(id uint) (*models.Category, error)
	FindBySlug(slug string) (*models.Category, error)
	FindAll(filter map[string]interface{}, page, limit int) ([]models.Category, int64, error)
	Update(category *models.Category) error
	Delete(id uint) error

	// Category-specific operations
	FindRootCategories() ([]models.Category, error)
	FindByParentID(parentID uint) ([]models.Category, error)
	FindFeaturedCategories(limit int) ([]models.Category, error)
	GetCategoryTree() ([]models.Category, error)
	GetProductCount(categoryID uint) (int64, error)
	CheckSlugExists(slug string, excludeID uint) (bool, error)
	GetBreadcrumb(categoryID uint) ([]models.Category, error)
	UpdateSortOrder(categoryID uint, sortOrder int) error
}

type categoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository creates a new instance of category repository
func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}
