package selleraccountsettingservice

import (
	"errors"
)

func (s *SellerAccountSettingService) UpdateCommission(sellerID uint, commission float64) error {
	if commission < 15 {
		return errors.New("commission must be at least 15%")
	}

	// Get current seller commission
	seller, err := s.repo.GetSellerByID(sellerID)
	if err != nil {
		return errors.New("seller not found")
	}

	// Commission can only be set during registration, not changed later
	if seller.Commission > 0 {
		return errors.New("commission cannot be changed after registration")
	}

	return s.repo.UpdateCommission(sellerID, commission)
}
