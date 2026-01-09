package cartservice

import (
	"fmt"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	cartitem "github.com/saadahmedbd/Treestore/Rest/DTO/CartItem"
)

// getCartWithBuyerAddress fetches cart with buyer's default address for shipping calculation
func (s *cartService) getCartWithBuyerAddress(buyerID uint) (*cartitem.CartSummaryResponse, error) {
	items, err := s.cartRepo.GetCart(buyerID, true)
	if err != nil {
		return nil, err
	}

	var address models.Address
	if err := Config.DB.Where("buyer_id = ? AND is_default = ?", buyerID, true).First(&address).Error; err == nil {
		addressStr := fmt.Sprintf("%s, %s, %s", address.City, address.District, address.State)
		return s.buildCartSummary(buyerID, items, addressStr)
	}
	return s.buildCartSummary(buyerID, items, "")
}
