package selleraccountservice

import (
	"fmt"

	selleraccount "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccount"
)

func (s *sellerRegistrationService) CompleteSellerProfile(userID uint, req *selleraccount.CompleteSellerProfileRequest) (*selleraccount.SellerProfileResponse, error) {
	seller, err := s.sellerRepo.GetSellerByUserId(userID)
	if err != nil {
		return nil, err
	}
	//update business info
	if req.BusinessEmail != "" {
		seller.BusinessEmail = req.BusinessEmail
	}
	if req.Phone != "" {
		seller.Phone = req.Phone
	}
	if req.StoreDesc != "" {
		seller.StoreDesc = req.StoreDesc
	}
	if req.BusinessType != "" {
		seller.BusinessType = req.BusinessType
	}

	// Update address
	seller.Address = req.Address
	seller.City = req.City
	seller.State = req.State
	seller.Country = req.Country
	if seller.Country == "" {
		seller.Country = "Bangladesh"
	}
	seller.PostalCode = req.PostalCode

	if err := s.sellerRepo.UpdateSellerProfile(seller); err != nil {
		return nil, fmt.Errorf("failed to update profile: %v", err)
	}

	return s.GetSellerProfile(userID)

}
