package categoryservice

import "errors"

// checkCircularReference prevents circular parent-child relationships
func (s *categoryService) checkCircularReference(categoryID, newParentID uint) error {
	// Get all ancestors of the new parent
	breadcrumb, err := s.categoryRepo.GetBreadcrumb(newParentID)
	if err != nil {
		return err
	}

	// Check if current category is in the ancestor chain
	for _, ancestor := range breadcrumb {
		if ancestor.ID == categoryID {
			return errors.New("circular reference detected: category cannot be a parent of its ancestor")
		}
	}

	return nil
}
