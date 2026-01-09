package selleraccountsettingservice

import (
	"errors"

	selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"
	util "github.com/saadahmedbd/Treestore/Util"
	"golang.org/x/crypto/bcrypt"
)

// ChangePassword changes seller password
func (s *SellerAccountSettingService) ChangePassword(sellerID uint, req *selleraccountsetting.ChangePasswordRequest) error {
	// Validate request
	if err := util.ValidateStruct(req); err != nil {
		return err
	}

	// Get seller and RegUser
	seller, err := s.repo.GetSellerByID(sellerID)
	if err != nil {
		return err
	}

	regUser, err := s.repo.GetRegUserByID(seller.UserId)
	if err != nil {
		return err
	}

	// Verify current password
	if err := bcrypt.CompareHashAndPassword([]byte(regUser.Password), []byte(req.CurrentPassword)); err != nil {
		return errors.New("current password is incorrect")
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	// Update password
	if err := s.repo.UpdatePassword(regUser.ID, string(hashedPassword)); err != nil {
		return errors.New("failed to update password")
	}

	return nil
}
