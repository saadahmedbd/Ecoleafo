package categoryservice

import models "github.com/saadahmedbd/Treestore/Models"

func (s *categoryService) GetProductsByCategory(categoryID uint, page, limit int) ([]models.Product, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	return s.productRepo.GetProductsByCategory(categoryID, page, limit)
}
