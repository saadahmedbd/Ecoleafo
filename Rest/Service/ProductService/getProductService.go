package productservice

import models "github.com/saadahmedbd/Treestore/Models"

func (s *ProductService) GetProduct(id uint) (*models.Product, error) {
	var product models.Product
	if err := s.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
