package productservice

import (
	"errors"
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (s *ProductService) GetProductWithRelations(productID uint) (*models.Product, error) {

	var product models.Product

	// Load product with all its relationships
	err := s.db.
		Preload("Seller").
		Preload("Category").
		Preload("Images").
		Preload("Attributes").
		Where("id = ?", productID).
		First(&product).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("product with ID %d not found", productID)
		}
		return nil, fmt.Errorf("failed to load product: %v", err)
	}

	return &product, nil
}
