package categoryservice

import categorydto "github.com/saadahmedbd/Treestore/Rest/DTO/CategoryDTO"

// GetAllCategories retrieves categories with filters and pagination
func (s *categoryService) GetAllCategories(filter categorydto.CategoryFilterRequest) ([]categorydto.CategoryResponse, int64, error) {
	// Set default pagination
	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.Limit < 1 {
		filter.Limit = 20
	}

	// Build filter map
	filterMap := make(map[string]interface{})
	if filter.ParentID != nil {
		filterMap["parent_id"] = filter.ParentID
	}
	if filter.IsFeatured != nil {
		filterMap["is_featured"] = *filter.IsFeatured
	}
	if filter.IsActive != nil {
		filterMap["is_active"] = *filter.IsActive
	}
	if filter.Search != "" {
		filterMap["search"] = filter.Search
	}

	// Get categories from repository
	categories, total, err := s.categoryRepo.FindAll(filterMap, filter.Page, filter.Limit)
	if err != nil {
		return nil, 0, err
	}

	// Map to response DTOs
	responses := make([]categorydto.CategoryResponse, len(categories))
	for i, cat := range categories {
		resp, _ := s.mapToCategoryResponse(&cat)
		responses[i] = *resp
	}

	return responses, total, nil
}
