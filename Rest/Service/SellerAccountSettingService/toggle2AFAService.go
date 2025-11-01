package selleraccountsettingservice

import (
	"errors"

	selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"
)

// Toggle2FA enables or disables 2FA
func (s *SellerAccountSettingService) Toggle2FA(sellerID uint, enabled bool) (*selleraccountsetting.Toggle2FAResponse, error) {
	// Get seller to find RegUser ID
	seller, err := s.repo.GetSellerByID(sellerID)
	if err != nil {
		return nil, err
	}

	// Toggle 2FA
	if err := s.repo.Toggle2FA(seller.UserId, enabled); err != nil {
		return nil, errors.New("failed to update 2FA settings")
	}

	message := "Two-factor authentication disabled successfully"
	if enabled {
		message = "Two-factor authentication enabled successfully"
	}

	return &selleraccountsetting.Toggle2FAResponse{
		Enabled: enabled,
		Message: message,
	}, nil
}
