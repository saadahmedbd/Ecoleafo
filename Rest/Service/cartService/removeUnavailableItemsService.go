package cartservice

func (s *cartService) RemoveUnavailableItems(buyerID uint) error {
	return s.cartRepo.RemoveUnavailableItems(buyerID)
}
