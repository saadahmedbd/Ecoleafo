package cartservice

func (s *cartService) GetCartItemCount(buyerID uint) (int, error) {
	return s.cartRepo.GetCartItemCount(buyerID)
}

func (s *cartService) GetWishlistCount(buyerID uint) (int, error) {
	return s.cartRepo.GetWishlistCount(buyerID)
}
