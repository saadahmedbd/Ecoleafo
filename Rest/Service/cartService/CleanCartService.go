package cartservice

func (s *cartService) ClearCart(buyerID uint) error {
	return s.cartRepo.ClearCart(buyerID)
}
