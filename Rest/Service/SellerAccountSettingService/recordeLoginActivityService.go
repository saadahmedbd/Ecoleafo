package selleraccountsettingservice

import models "github.com/saadahmedbd/Treestore/Models"

// RecordLoginActivity records a new login activity
func (s *SellerAccountSettingService) RecordLoginActivity(sellerID uint, device, browser, os, location, ipAddress string) error {
	activity := &models.SellerLoginActivity{
		SellerID:  sellerID,
		Device:    device,
		Browser:   browser,
		OS:        os,
		Location:  location,
		IPAddress: ipAddress,
		IsCurrent: true,
	}

	if err := s.repo.CreateLoginActivity(activity); err != nil {
		return err
	}

	// Set this as current and unmark others
	if err := s.repo.SetCurrentLoginActivity(sellerID, activity.ID); err != nil {
		return err
	}

	return nil
}
