package cartservice

func (s *cartService) DecrementQuantity(buyerID uint, productID uint) error {
	return s.cartRepo.DecrementQuantity(buyerID, productID)
}
