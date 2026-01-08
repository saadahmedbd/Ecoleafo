package categoryservice

import (
	"errors"

	categorydto "github.com/saadahmedbd/Treestore/Rest/DTO/CategoryDTO"
)

// UpdateCategoryIcon updates the icon URL for a category
func (s *categoryService) UpdateCategoryIcon(id uint, iconURL string) (*categorydto.CategoryResponse, error) {
	// Get existing category
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("category not found")
	}

	// Update icon
	category.Icon = iconURL
	if err := s.categoryRepo.Update(category); err != nil {
		return nil, err
	}

	// Map to response
	return s.mapToCategoryResponse(category)
}
