package cartservice

import cartitem "github.com/saadahmedbd/Treestore/Rest/DTO/CartItem"

func (s *cartService) GetCartSummary(buyerID uint) (*cartitem.CartSummaryResponse, error) {
	return s.getCartWithBuyerAddress(buyerID)
}
