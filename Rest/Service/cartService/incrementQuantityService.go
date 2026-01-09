package cartservice

import (
	"fmt"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
)

func (s *cartService) IncrementQuantity(buyerID uint, productID uint) error {
	// Validate stock before increment
	item, err := s.cartRepo.GetCartItem(buyerID, productID)
	if err != nil {
		return err
	}

	var product models.Product
	if err := Config.DB.Where("id = ?", productID).First(&product).Error; err != nil {
		return err
	}

	if item.Quantity >= product.Quantity {
		return fmt.Errorf("maximum quantity reached")
	}

	return s.cartRepo.IncrementQuantity(buyerID, productID)
}
