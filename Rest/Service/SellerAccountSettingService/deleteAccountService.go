package selleraccountsettingservice

import (
	"errors"

	selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"
	"golang.org/x/crypto/bcrypt"
)

// DeleteAccount permanently deletes seller account
func (s *SellerAccountSettingService) DeleteAccount(sellerID uint, password string) (*selleraccountsetting.DeleteAccountResponse, error) {
	// Get seller and RegUser
	seller, err := s.repo.GetSellerByID(sellerID)
	if err != nil {
		return nil, err
	}

	regUser, err := s.repo.GetRegUserByID(seller.UserId)
	if err != nil {
		return nil, err
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(regUser.Password), []byte(password)); err != nil {
		return nil, errors.New("incorrect password")
	}

	// Soft delete seller
	if err := s.repo.DeleteSeller(sellerID); err != nil {
		return nil, errors.New("failed to delete seller account")
	}

	// Soft delete RegUser
	if err := s.repo.DeleteRegUser(regUser.ID); err != nil {
		return nil, errors.New("failed to delete user account")
	}

	return &selleraccountsetting.DeleteAccountResponse{
		Message: "Account deleted successfully",
		Success: true,
	}, nil
}
