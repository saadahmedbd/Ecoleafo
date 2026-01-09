package selleraccountsettingservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"
)

// GetFullProfile retrieves complete seller profile
func (s *SellerAccountSettingService) GetFullProfile(userID uint) (*selleraccountsetting.SellerProfileResponse, error) {
	// Get seller data by RegUser ID (from JWT)
	seller, err := s.repo.GetSellerByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Get RegUser data
	regUser, err := s.repo.GetRegUserByID(seller.UserId)
	if err != nil {
		return nil, err
	}

	// Get policies using actual seller ID
	policies, err := s.repo.GetSellerPolicies(seller.ID)
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
