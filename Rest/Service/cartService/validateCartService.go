package cartservice

import (
	"fmt"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	cartitem "github.com/saadahmedbd/Treestore/Rest/DTO/CartItem"
)

func (s *cartService) ValidateCart(buyerID uint) (*cartitem.CartSummaryResponse, error) {
	// Get only selected items
	selectedItems, err := s.cartRepo.GetSelectedCartItems(buyerID)
	if err != nil {
		return nil, err
	}

	// Get buyer's address
	var address models.Address
	if err := Config.DB.Where("buyer_id = ? AND is_default = ?", buyerID, true).First(&address).Error; err == nil {
		addressStr := fmt.Sprintf("%s, %s, %s", address.City, address.District, address.State)
		return s.buildCartSummary(buyerID, selectedItems, addressStr)
	}

	return s.buildCartSummary(buyerID, selectedItems, "")
}

func (s *cartService) ValidateCartWithAddress(buyerID uint, address string) (*cartitem.CartSummaryResponse, error) {
	// Get only selected items
	selectedItems, err := s.cartRepo.GetSelectedCartItems(buyerID)
	if err != nil {
		return nil, err
	}

	return s.buildCartSummary(buyerID, selectedItems, address)
}
