package categoryservice

import categorydto "github.com/saadahmedbd/Treestore/Rest/DTO/CategoryDTO"

// GetCategoryByID retrieves a category by ID
func (s *categoryService) GetCategoryByID(id uint) (*categorydto.CategoryResponse, error) {
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return s.mapToCategoryResponse(category)
}
