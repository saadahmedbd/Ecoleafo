package cartservice

import (
	"fmt"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	cartitem "github.com/saadahmedbd/Treestore/Rest/DTO/CartItem"
)

func (s *cartService) UpdateCartItem(buyerID uint, productID uint, req cartitem.UpdateCartItemRequest) error {
	item, err := s.cartRepo.GetCartItem(buyerID, productID)
	if err != nil {
		return err
	}
	// Validate stock
	var product models.Product
	if err := Config.DB.Where("id = ?", productID).First(&product).Error; err != nil {
		return fmt.Errorf("product not found")
	}

	if product.Quantity < req.Quantity {
		return fmt.Errorf("insufficient stock")
	}

	item.Quantity = req.Quantity
	item.IsGift = req.IsGift
	item.GiftMessage = req.GiftMessage

	return s.cartRepo.UpdateCartItem(item)
}
