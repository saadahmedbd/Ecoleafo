package cartservice

func (s *cartService) RemoveFromCart(buyerID uint, productID uint) error {
	return s.cartRepo.RemoveFromCart(buyerID, productID)
}
