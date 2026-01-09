package categoryservice

import categorydto "github.com/saadahmedbd/Treestore/Rest/DTO/CategoryDTO"

// GetRootCategories retrieves all top-level categories
func (s *categoryService) GetRootCategories() ([]categorydto.CategoryResponse, error) {
	categories, err := s.categoryRepo.FindRootCategories()
	if err != nil {
		return nil, err
	}

	responses := make([]categorydto.CategoryResponse, len(categories))
	for i, cat := range categories {
		resp, _ := s.mapToCategoryResponse(&cat)
		responses[i] = *resp
	}

	return responses, nil
}
