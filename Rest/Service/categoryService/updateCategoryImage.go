package categoryservice

import (
	"errors"

	categorydto "github.com/saadahmedbd/Treestore/Rest/DTO/CategoryDTO"
)

// UpdateCategoryImage updates the image URL for a category
func (s *categoryService) UpdateCategoryImage(id uint, imageURL string) (*categorydto.CategoryResponse, error) {
	// Get existing category
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("category not found")
	}

	// Update image
	category.Image = imageURL
	if err := s.categoryRepo.Update(category); err != nil {
		return nil, err
	}

	// Map to response
	return s.mapToCategoryResponse(category)
}
