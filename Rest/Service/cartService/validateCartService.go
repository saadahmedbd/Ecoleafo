package cartservice

import cartitem "github.com/saadahmedbd/Treestore/Rest/DTO/CartItem"

func (s *cartService) ValidateCart(buyerID uint) (*cartitem.CartSummaryResponse, error) {
	invalidIDs, err := s.cartRepo.ValidateCartStock(buyerID)
	if err != nil {
		return nil, err
	}

	if len(invalidIDs) > 0 {
		// Mark items as unavailable or remove them
		// For now, just return the cart with validation info
	}

	return s.GetCart(buyerID)
}
