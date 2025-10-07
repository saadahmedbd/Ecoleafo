package selleraccountservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	selleraccount "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccount"
)

func (s *sellerRegistrationService) AddPaymentMethod(userID uint, req *selleraccount.AddPaymentMethodRequest) (*selleraccount.PaymentMethodInfo, error) {
	seller, err := s.sellerRepo.GetSellerByUserId(userID)
	if err != nil {
		return nil, err
	}
	paymentMethod := &models.SellerPaymentMethod{
		SellerID:      seller.ID,
		Type:          req.Type,
		AccountName:   req.AccountName,
		AccountNumber: req.AccountNumber,
		BankName:      req.BankName,
		BankCode:      req.BankCode,
		RoutingNumber: req.RoutingNumber,
		IsDefault:     req.IsDefault,
		IsActive:      true,
	}

	if err := s.sellerRepo.CreatePaymentMethod(paymentMethod); err != nil {
		return nil, err
	}
	return &selleraccount.PaymentMethodInfo{
		ID:            paymentMethod.ID,
		Type:          paymentMethod.Type,
		AccountName:   paymentMethod.AccountName,
		AccountNumber: paymentMethod.AccountNumber,
		BankName:      paymentMethod.BankName,
		IsDefault:     paymentMethod.IsDefault,
		IsActive:      paymentMethod.IsActive,
	}, nil
}
