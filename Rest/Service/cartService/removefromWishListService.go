package cartservice

func (s *cartService) RemoveFromWishlist(buyerID uint, productID uint) error {
	return s.cartRepo.RemoveFromWishlist(buyerID, productID)
}
