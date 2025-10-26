package categoryservice

import (
	"errors"

	categorydto "github.com/saadahmedbd/Treestore/Rest/DTO/CategoryDTO"
	util "github.com/saadahmedbd/Treestore/Util"
)

// UpdateCategory updates an existing category
func (s *categoryService) UpdateCategory(id uint, req categorydto.UpdateCategoryRequest) (*categorydto.CategoryResponse, error) {
	// Get existing category
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Name != nil {
		category.Name = *req.Name
		// Regenerate slug if name changed
		newSlug := util.GenerateSlug(*req.Name)
		exists, _ := s.categoryRepo.CheckSlugExists(newSlug, category.ID)
		if !exists {
			category.Slug = newSlug
		}
	}
	if req.Description != nil {
		category.Description = *req.Description
	}
	if req.Image != nil {
		category.Image = *req.Image
	}
	if req.Icon != nil {
		category.Icon = *req.Icon
	}
	if req.ParentID != nil {
		// Validate parent
		if *req.ParentID == category.ID {
			return nil, errors.New("category cannot be its own parent")
		}
		// Check for circular reference
		if err := s.checkCircularReference(category.ID, *req.ParentID); err != nil {
			return nil, err
		}
		category.ParentID = req.ParentID
	}
	if req.SortOrder != nil {
		category.SortOrder = *req.SortOrder
	}
	if req.IsFeatured != nil {
		category.IsFeatured = *req.IsFeatured
	}
	if req.IsActive != nil {
		category.IsActive = *req.IsActive
	}
	if req.MetaTitle != nil {
		category.MetaTitle = *req.MetaTitle
	}
	if req.MetaDescription != nil {
		category.MetaDescription = *req.MetaDescription
	}
	if req.MetaKeywords != nil {
		category.MetaKeywords = *req.MetaKeywords
	}

	// Save updates
	if err := s.categoryRepo.Update(category); err != nil {
		return nil, err
	}

	return s.mapToCategoryResponse(category)
}
