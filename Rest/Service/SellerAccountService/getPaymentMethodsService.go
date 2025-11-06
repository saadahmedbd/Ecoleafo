package selleraccountservice

import selleraccount "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccount"

func (s *sellerRegistrationService) GetPaymentMethods(userID uint) ([]selleraccount.PaymentMethodInfo, error) {
	seller, err := s.sellerRepo.GetSellerByRegUserId(userID)
	if err != nil {
		return nil, err
	}
	paymentMethods, err := s.sellerRepo.GetPaymentMethods(seller.ID)
	if err != nil {
		return nil, err
	}
	var paymentMethodInfos []selleraccount.PaymentMethodInfo
	for _, pm := range paymentMethods {
		paymentMethodInfos = append(paymentMethodInfos, selleraccount.PaymentMethodInfo{
			ID:            pm.ID,
			Type:          pm.Type,
			AccountName:   pm.AccountName,
			AccountNumber: pm.AccountNumber,
			BankName:      pm.BankName,
			IsDefault:     pm.IsDefault,
			IsActive:      pm.IsActive,
		})

	}
	return paymentMethodInfos, nil
}
