package selleraccountsettingservice

import (
	"errors"

	selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"
	util "github.com/saadahmedbd/Treestore/Util"
)

// UpdateAccount updates seller account information
func (s *SellerAccountSettingService) UpdateAccount(sellerID uint, req *selleraccountsetting.UpdateAccountRequest) error {
	// Validate request
	if err := util.ValidateStruct(req); err != nil {
		return err
	}

	// Get seller to find RegUser ID
	seller, err := s.repo.GetSellerByID(sellerID)
	if err != nil {
		return err
	}

	// Prepare updates
	updates := map[string]interface{}{
		"first_name": req.FirstName,
		"last_name":  req.LastName,
	}

	// Update RegUser
	if err := s.repo.UpdateRegUser(seller.UserId, updates); err != nil {
		return errors.New("failed to update account information")
	}

	// Update seller phone
	sellerUpdates := map[string]interface{}{
		"phone": req.Phone,
	}
	if err := s.repo.UpdateSeller(sellerID, sellerUpdates); err != nil {
		return errors.New("failed to update phone number")
	}

	return nil
}
