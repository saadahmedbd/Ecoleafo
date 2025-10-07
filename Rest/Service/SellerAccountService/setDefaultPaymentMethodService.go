package selleraccountservice

func (s *sellerRegistrationService) SetDefaultPaymentMethod(userID uint, methodID uint) error {
	seller, err := s.sellerRepo.GetSellerByUserId(userID)
	if err != nil {
		return err
	}

	return s.sellerRepo.SetDefaultPaymentMethod(seller.ID, methodID)
}
