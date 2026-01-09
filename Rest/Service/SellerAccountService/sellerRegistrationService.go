package selleraccountservice

import (
	selleraccount "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccount"
	selleraccountrepo "github.com/saadahmedbd/Treestore/Rest/Repository/SellerAccountRepo"
)

type SellerRegistrationService interface {
	//registration
	RegisterSeller(req *selleraccount.SellerRegistationRequest) (*selleraccount.SellerRegistrationResponse, error)
	CheckProfileStatus(userID uint) (*selleraccount.SellerProfileCompletionStatus, error)
	CompleteSellerProfile(userID uint, req *selleraccount.CompleteSellerProfileRequest) (*selleraccount.SellerProfileResponse, error)

	//profile mangment
	GetSellerProfile(userID uint) (*selleraccount.SellerProfileResponse, error)
	UpdateStoreInfo(userID uint, req *selleraccount.UpdateStoreInfoRequest) error

	//payment method
	AddPaymentMethod(UserID uint, req *selleraccount.AddPaymentMethodRequest) (*selleraccount.PaymentMethodInfo, error)
	GetPaymentMethods(userID uint) ([]selleraccount.PaymentMethodInfo, error)
	UpdatePaymentMethod(userID uint, paymentMethodID uint, req *selleraccount.UpdatePaymentMethodRequest) error
	DeletePaymentMethod(userID uint, paymentMethodID uint) error
	SetDefaultPaymentMethod(userID uint, paymentMethodID uint) error
}

type sellerRegistrationService struct {
	sellerRepo selleraccountrepo.SellerRegistrationRepository
}

func NewSellerRegistrationService(sellerRepo selleraccountrepo.SellerRegistrationRepository) SellerRegistrationService {
	return &sellerRegistrationService{
		sellerRepo: sellerRepo,
	}
}
