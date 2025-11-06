package selleraccountservice

func (s *sellerRegistrationService) SetDefaultPaymentMethod(userID uint, methodID uint) error {
	seller, err := s.sellerRepo.GetSellerByRegUserId(userID)
	if err != nil {
		return err
	}

	return s.sellerRepo.SetDefaultPaymentMethod(seller.ID, methodID)
}
