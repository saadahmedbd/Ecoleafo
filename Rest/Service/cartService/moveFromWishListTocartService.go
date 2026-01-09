package cartservice

func (s *cartService) MoveFromWishlistToCart(buyerID uint, productID uint) error {
	return s.cartRepo.MoveFromWishlistToCart(buyerID, productID, 1)
}
