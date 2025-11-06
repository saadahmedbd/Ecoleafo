package selleraccountservice

func (s *sellerRegistrationService) DeletePaymentMethod(userID uint, paymentMethodID uint) error {
	seller, err := s.sellerRepo.GetSellerByRegUserId(userID)
	if err != nil {
		return err
	}

	return s.sellerRepo.DeletePaymentMethod(paymentMethodID, seller.ID)
}
