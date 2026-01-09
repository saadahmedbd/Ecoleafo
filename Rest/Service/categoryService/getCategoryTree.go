package categoryservice

import categorydto "github.com/saadahmedbd/Treestore/Rest/DTO/CategoryDTO"

// GetCategoryTree builds hierarchical tree for navigation menus
func (s *categoryService) GetCategoryTree() ([]categorydto.CategoryTreeResponse, error) {
	categories, err := s.categoryRepo.GetCategoryTree()
	if err != nil {
		return nil, err
	}

	// Build tree structure
	return s.buildTree(categories, nil), nil
}
