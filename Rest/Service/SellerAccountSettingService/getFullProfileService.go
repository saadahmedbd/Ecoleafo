package selleraccountsettingservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"
)

// GetFullProfile retrieves complete seller profile
func (s *SellerAccountSettingService) GetFullProfile(sellerID uint) (*selleraccountsetting.SellerProfileResponse, error) {
	// Get seller data
	seller, err := s.repo.GetSellerByID(sellerID)
	if err != nil {
		return nil, err
	}

	// Get RegUser data
	regUser, err := s.repo.GetRegUserByID(seller.UserId)
	if err != nil {
		return nil, err
	}

	// Get policies
	policies, err := s.repo.GetSellerPolicies(sellerID)
	if err != nil {
		// Don't fail if policies not found, just continue
		policies = &models.SellerPolicy{}
	}

	// Convert to DTO
	response := selleraccountsetting.ToSellerProfileResponse(seller, regUser)
	response.ReturnPolicy = policies.ReturnPolicy
	response.ShippingPolicy = policies.ShippingPolicy
	response.FAQ = policies.FAQ

	// Get 2FA status
	twoFactorEnabled, _ := s.repo.Get2FAStatus(regUser.ID)
	response.TwoFactorEnabled = twoFactorEnabled

	return response, nil
}
