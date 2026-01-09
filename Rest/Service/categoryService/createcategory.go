package categoryservice

import (
	"errors"
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	categorydto "github.com/saadahmedbd/Treestore/Rest/DTO/CategoryDTO"
	util "github.com/saadahmedbd/Treestore/Util"
)

// CreateCategory creates a new category with validation
func (s *categoryService) CreateCategory(req categorydto.CreateCategoryRequest) (*categorydto.CategoryResponse, error) {
	// Generate slug from name (e.g., "Smart Phones" -> "smart-phones")
	slug := util.GenerateSlug(req.Name)

	// Check if slug already exists
	exists, err := s.categoryRepo.CheckSlugExists(slug, 0)
	if err != nil {
		return nil, err
	}
	if exists {
		// Append a number to make it unique
		slug = fmt.Sprintf("%s-%d", slug, util.GenerateRandomNumber())
	}

	// Validate parent category if provided
	if req.ParentID != nil {
		parent, err := s.categoryRepo.FindByID(*req.ParentID)
		if err != nil {
			return nil, errors.New("parent category not found")
		}
		if !parent.IsActive {
			return nil, errors.New("parent category is not active")
		}
	}

	// Create category model
	category := &models.Category{
		Name:        req.Name,
		Slug:        slug,
		Description: req.Description,

		ParentID:        req.ParentID,
		SortOrder:       req.SortOrder,
		IsFeatured:      req.IsFeatured,
		IsActive:        req.IsActive,
		MetaTitle:       req.MetaTitle,
		MetaDescription: req.MetaDescription,
		MetaKeywords:    req.MetaKeywords,
	}

	// Save to database
	if err := s.categoryRepo.Create(category); err != nil {
		return nil, err
	}

	// Return response
	return s.mapToCategoryResponse(category)
}
