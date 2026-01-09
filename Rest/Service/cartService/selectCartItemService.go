package cartservice

func (s *cartService) SelectCartItem(buyerID uint, productID uint, isSelected bool) error {
	return s.cartRepo.UpdateCartItemSelection(buyerID, productID, isSelected)
}

func (s *cartService) SelectCartItemByID(buyerID uint, cartItemID uint, isSelected bool) error {
	return s.cartRepo.UpdateCartItemSelectionByID(buyerID, cartItemID, isSelected)
}

func (s *cartService) SelectAllCartItems(buyerID uint, isSelected bool) error {
	return s.cartRepo.UpdateAllCartItemsSelection(buyerID, isSelected)
}
