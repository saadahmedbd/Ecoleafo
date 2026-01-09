package guestcartservice

func (s *guestcartservice) MigrateGuestCart(sessionID string, userID uint) error {
	buyer, err := s.buyerRepo.GetBuyerByUserID(userID)
	if err != nil {
		return err
	}

	return s.cartRepo.MigrateGuestCartToUser(sessionID, buyer.ID)
}
