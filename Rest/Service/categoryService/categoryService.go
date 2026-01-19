package categoryservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	categorydto "github.com/saadahmedbd/Treestore/Rest/DTO/CategoryDTO"
	categoryrepo "github.com/saadahmedbd/Treestore/Rest/Repository/CategoryRepo"
	productrepo "github.com/saadahmedbd/Treestore/Rest/Repository/ProductRepo"
)

type CategoryService interface {
	// CRUD operations
	CreateCategory(req categorydto.CreateCategoryRequest) (*categorydto.CategoryResponse, error)
	GetCategoryByID(id uint) (*categorydto.CategoryResponse, error)
	GetCategoryBySlug(slug string) (*categorydto.CategoryResponse, error)
	GetAllCategories(filter categorydto.CategoryFilterRequest) ([]categorydto.CategoryResponse, int64, error)
	GetCategoriesForSeller() ([]categorydto.CategoryResponse, error)
	UpdateCategory(id uint, req categorydto.UpdateCategoryRequest) (*categorydto.CategoryResponse, error)
	UpdateCategoryImage(id uint, imageURL string) (*categorydto.CategoryResponse, error)
	UpdateCategoryIcon(id uint, iconURL string) (*categorydto.CategoryResponse, error)
	DeleteCategory(id uint) error

	// Category-specific operations
	GetRootCategories() ([]categorydto.CategoryResponse, error)
	GetSubcategories(parentID uint) ([]categorydto.CategoryResponse, error)
	GetFeaturedCategories(limit int) ([]categorydto.CategoryResponse, error)
	GetCategoryTree() ([]categorydto.CategoryTreeResponse, error)
	GetBreadcrumb(categoryID uint) ([]categorydto.CategoryBriefResponse, error)
	ReorderCategories(updates []struct {
		ID        uint
		SortOrder int
	}) error

	// Product operations
	GetProductsByCategory(categoryID uint, page, limit int) ([]models.Product, int64, error)
}

type categoryService struct {
	categoryRepo categoryrepo.CategoryRepository
	productRepo  productrepo.ProductRepository
}

// NewCategoryService creates a new instance of category service
func NewCategoryService(
	categoryRepo categoryrepo.CategoryRepository,
	productRepo productrepo.ProductRepository,
) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
		productRepo:  productRepo,
	}
}
