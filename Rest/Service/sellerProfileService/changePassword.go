package sellerprofileservice

import (
	"fmt"

	sellerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/SellerProfile"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (s *sellerService) ChangePassword(userIDFromJWT uint, req sellerprofile.ChangePasswordRequest) error {
	// Validate passwords match
	if req.NewPassword != req.ConfirmPassword {
		return fmt.Errorf("passwords do not match")
	}

	// First get seller by RegUser ID
	seller, err := s.sellerRepo.GetSellerByUserID(userIDFromJWT)
	if err != nil {
		return err
	}

	// Then get full seller with relations
	sellerWithRelations, err := s.sellerRepo.GetSellerWithRelations(seller.ID)
	if err != nil {
		return err
	}

	// Verify current password using RegUser password
	if err := util.CheckPassword(sellerWithRelations.RegUser.Password, req.CurrentPassword); err != nil {
		return fmt.Errorf("current password is incorrect")
	}

	// Hash new password
	hashedPassword, err := util.HashPassword(req.NewPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	// Update RegUser password directly in database
	sellerWithRelations.RegUser.Password = hashedPassword
	return s.sellerRepo.UpdateRegUser(sellerWithRelations.RegUser)
}
