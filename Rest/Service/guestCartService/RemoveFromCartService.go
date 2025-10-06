package guestcartservice

func (s *guestcartservice) RemoveCartItem(sessionID string, userID *uint, productID uint) error {
	if userID != nil {
		buyer, err := s.buyerRepo.GetBuyerByUserID(*userID)
		if err != nil {
			// If buyer doesn't exist, treat as guest
			if err.Error() == "buyer not found" {
				return s.cartRepo.RemoveGuestCartItem(sessionID, productID)
			}
			return err
		}
		return s.cartRepo.RemoveCartItem(buyer.ID, productID)
	}
	return s.cartRepo.RemoveGuestCartItem(sessionID, productID)
}
