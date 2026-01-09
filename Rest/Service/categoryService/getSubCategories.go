package categoryservice

import categorydto "github.com/saadahmedbd/Treestore/Rest/DTO/CategoryDTO"

// GetSubcategories retrieves child categories
func (s *categoryService) GetSubcategories(parentID uint) ([]categorydto.CategoryResponse, error) {
	categories, err := s.categoryRepo.FindByParentID(parentID)
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
