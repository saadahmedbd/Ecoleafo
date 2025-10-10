package cartservice

func (s *cartService) GetCartItemCount(buyerID uint) (int, error) {
	return s.cartRepo.GetCartItemCount(buyerID)
}
