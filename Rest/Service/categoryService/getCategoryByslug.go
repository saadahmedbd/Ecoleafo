package categoryservice

import categorydto "github.com/saadahmedbd/Treestore/Rest/DTO/CategoryDTO"

// GetCategoryBySlug retrieves a category by slug (for SEO-friendly URLs)
func (s *categoryService) GetCategoryBySlug(slug string) (*categorydto.CategoryResponse, error) {
	category, err := s.categoryRepo.FindBySlug(slug)
	if err != nil {
		return nil, err
	}

	return s.mapToCategoryResponse(category)
}
