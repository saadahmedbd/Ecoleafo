package commissionpayoutearningservice

func (s *commissionService) GetSellerIDByUserID(userID uint) (uint, error) {
	seller, err := s.sellerRepo.GetSellerByUserID(userID)
	if err != nil {
		return 0, err
	}
	return seller.ID, nil
}
