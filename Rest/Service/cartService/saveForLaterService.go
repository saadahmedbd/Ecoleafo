package cartservice

func (s *cartService) SaveForLater(buyerID uint, productID uint) error {
	return s.cartRepo.SaveForLater(buyerID, productID)
}
