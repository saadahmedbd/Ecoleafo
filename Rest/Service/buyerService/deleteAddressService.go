package buyerservice

func (s *buyerService) DeleteAddress(userIDFromJWT uint, addressID uint) error {
	buyer, err := s.buyerRepo.GetBuyerByUserID(userIDFromJWT)
	if err != nil {
		return err
	}

	return s.buyerRepo.DeleteAddress(addressID, buyer.ID)
}
