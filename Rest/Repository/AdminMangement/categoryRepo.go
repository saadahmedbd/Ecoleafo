package adminmangement

import (
	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

// GetAll retrieves all categories
func (r *CategoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Preload("SubCategories").Where("parent_id IS NULL").Order("display_order ASC").Find(&categories).Error
	return categories, err
}

// GetByID retrieves a category by ID
func (r *CategoryRepository) GetByID(id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.Preload("Parent").Preload("SubCategories").First(&category, id).Error
	return &category, err
}

// GetBySlug retrieves a category by slug
func (r *CategoryRepository) GetBySlug(slug string) (*models.Category, error) {
	var category models.Category
	err := r.db.Where("slug = ?", slug).First(&category).Error
	return &category, err
}

// Create creates a new category
func (r *CategoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

// Update updates a category
func (r *CategoryRepository) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

// Delete soft deletes a category
func (r *CategoryRepository) Delete(id uint) error {
	return r.db.Delete(&models.Category{}, id).Error
}

// GetActiveCategories retrieves all active categories
func (r *CategoryRepository) GetActiveCategories() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Where("is_active = ?", true).Order("display_order ASC").Find(&categories).Error
	return categories, err
}
