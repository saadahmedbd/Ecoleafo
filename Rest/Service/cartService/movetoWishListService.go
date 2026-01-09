package cartservice

func (s *cartService) MoveToWishlist(buyerID uint, productID uint) error {
	return s.cartRepo.MoveToWishlist(buyerID, productID)
}
