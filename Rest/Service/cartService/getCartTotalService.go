package cartservice

func (s *cartService) GetCartTotal(buyerID uint) (float64, error) {
	return s.cartRepo.GetCartTotal(buyerID)
}
