package selleraccountservice

import (
	"fmt"

	selleraccount "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccount"
)

func (s *sellerRegistrationService) CompleteSellerProfile(userID uint, req *selleraccount.CompleteSellerProfileRequest) (*selleraccount.SellerProfileResponse, error) {
	seller, err := s.sellerRepo.GetSellerByRegUserId(userID)
	if err != nil {
		return nil, err
	}
	
	// Update business info
	if req.BusinessEmail != "" {
		seller.BusinessEmail = req.BusinessEmail
	}
	if req.Phone != "" {
		seller.Phone = req.Phone
	}
	seller.StoreDesc = req.StoreDesc
	seller.BusinessType = req.BusinessType

	// Update address - always update these fields
	seller.Address = req.Address
	seller.City = req.City
	seller.State = req.State
	seller.PostalCode = req.PostalCode
	if req.Country != "" {
		seller.Country = req.Country
	} else {
		seller.Country = "Bangladesh"
	}

	if err := s.sellerRepo.UpdateSellerProfile(seller); err != nil {
		return nil, fmt.Errorf("failed to update profile: %v", err)
	}

	return s.GetSellerProfile(userID)
}
