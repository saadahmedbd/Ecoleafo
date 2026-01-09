package categoryservice

import (
	"errors"
	"fmt"
)

// DeleteCategory soft deletes a category
func (s *categoryService) DeleteCategory(id uint) error {
	// Check if category exists
	_, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return err
	}

	// Check if category has children
	children, err := s.categoryRepo.FindByParentID(id)
	if err != nil {
		return err
	}
	if len(children) > 0 {
		return errors.New("cannot delete category with subcategories")
	}

	// Check if category has products
	productCount, err := s.categoryRepo.GetProductCount(id)
	if err != nil {
		return err
	}
	if productCount > 0 {
		return errors.New(fmt.Sprintf("cannot delete category with %d products", productCount))
	}

	// Delete category
	return s.categoryRepo.Delete(id)
}
