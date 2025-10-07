package selleraccountservice

import selleraccount "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccount"

func (s *sellerRegistrationService) UpdatePaymentMethod(userID uint, paymentMethodID uint, req *selleraccount.UpdatePaymentMethodRequest) error {
	seller, err := s.sellerRepo.GetSellerByUserId(userID)
	if err != nil {
		return err
	}
	paymentMethod, err := s.sellerRepo.GetPaymentMethodByID(paymentMethodID, seller.ID)
	if err != nil {
		return err
	}

	if req.AccountName != "" {
		paymentMethod.AccountName = req.AccountName
	}
	if req.AccountNumber != "" {
		paymentMethod.AccountNumber = req.AccountNumber
	}
	if req.BankName != "" {
		paymentMethod.BankName = req.BankName
	}
	if req.BankCode != "" {
		paymentMethod.BankCode = req.BankCode
	}
	if req.RoutingNumber != "" {
		paymentMethod.RoutingNumber = req.RoutingNumber
	}
	paymentMethod.IsActive = req.IsActive
	paymentMethod.IsDefault = req.IsDefault

	return s.sellerRepo.UpdatePaymentMethod(paymentMethod)

}
