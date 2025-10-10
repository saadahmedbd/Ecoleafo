package cartservice

func (s *cartService) MoveToCart(buyerID uint, productID uint) error {
	return s.cartRepo.MoveToCart(buyerID, productID)
}
