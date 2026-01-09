package selleraccountsettingservice

import (
	"errors"

	selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"
)

// DeactivateAccount deactivates seller account
func (s *SellerAccountSettingService) DeactivateAccount(sellerID uint) (*selleraccountsetting.DeactivateAccountResponse, error) {
	if err := s.repo.DeactivateSeller(sellerID); err != nil {
		return nil, errors.New("failed to deactivate account")
	}

	return &selleraccountsetting.DeactivateAccountResponse{
		Message: "Account deactivated successfully",
		Success: true,
	}, nil
}
