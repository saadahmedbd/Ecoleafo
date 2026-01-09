package categoryservice

import categorydto "github.com/saadahmedbd/Treestore/Rest/DTO/CategoryDTO"

// GetBreadcrumb retrieves category breadcrumb path
func (s *categoryService) GetBreadcrumb(categoryID uint) ([]categorydto.CategoryBriefResponse, error) {
	categories, err := s.categoryRepo.GetBreadcrumb(categoryID)
	if err != nil {
		return nil, err
	}

	breadcrumb := make([]categorydto.CategoryBriefResponse, len(categories))
	for i, cat := range categories {
		breadcrumb[i] = categorydto.CategoryBriefResponse{
			ID:   cat.ID,
			Name: cat.Name,
			Slug: cat.Slug,
			Icon: cat.Icon,
		}
	}

	return breadcrumb, nil
}
