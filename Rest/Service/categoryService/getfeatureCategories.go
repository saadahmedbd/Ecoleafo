package categoryservice

import categorydto "github.com/saadahmedbd/Treestore/Rest/DTO/CategoryDTO"

// GetFeaturedCategories retrieves featured categories for homepage
func (s *categoryService) GetFeaturedCategories(limit int) ([]categorydto.CategoryResponse, error) {
	categories, err := s.categoryRepo.FindFeaturedCategories(limit)
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
