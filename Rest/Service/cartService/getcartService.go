package cartservice

import cartitem "github.com/saadahmedbd/Treestore/Rest/DTO/CartItem"

func (s *cartService) GetCart(buyerID uint) (*cartitem.CartSummaryResponse, error) {
	items, err := s.cartRepo.GetCart(buyerID, true)
	if err != nil {
		return nil, err
	}
	return s.buildCartSummary(buyerID, items, "")
}

func (s *cartService) GetCartWithAddress(buyerID uint, address string) (*cartitem.CartSummaryResponse, error) {
	items, err := s.cartRepo.GetCart(buyerID, true)
	if err != nil {
		return nil, err
	}

	return s.buildCartSummary(buyerID, items, address)
}
