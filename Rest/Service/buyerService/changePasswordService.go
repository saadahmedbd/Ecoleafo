package buyerservice

import (
	"fmt"

	buyerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/BuyerProfile"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (s *buyerService) ChangePassword(userIDFromJWT uint, req buyerprofile.ChangeBuyerPassword) error {
	// Validate passwords match
	if req.NewPassword != req.ConfirmPassword {
		return fmt.Errorf("passwords do not match")
	}

	// First get buyer by RegUser ID
	buyer, err := s.buyerRepo.GetBuyerByUserID(userIDFromJWT)
	if err != nil {
		return err
	}

	// Then get full buyer with relations
	buyerWithRelations, err := s.buyerRepo.GetBuyerWithRelations(buyer.ID)
	if err != nil {
		return err
	}

	// Verify current password using RegUser password
	if err := util.CheckPassword(buyerWithRelations.RegUser.Password, req.CurrentPassword); err != nil {
		return fmt.Errorf("current password is incorrect")
	}

	// Hash new password
	hashedPassword, err := util.HashPassword(req.NewPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	// Update RegUser password directly in database
	buyerWithRelations.RegUser.Password = hashedPassword
	return s.buyerRepo.UpdateRegUser(buyerWithRelations.RegUser)

}
