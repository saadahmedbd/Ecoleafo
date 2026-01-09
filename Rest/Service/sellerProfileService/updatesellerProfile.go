package sellerprofileservice

import (
	"fmt"

	sellerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/SellerProfile"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (s *sellerService) UpdateSellerProfile(userIDFromJWT uint, req sellerprofile.UpdateSellerProfileRequest) (*sellerprofile.SellerProfileResponse, error) {
	// Get seller
	seller, err := s.sellerRepo.GetSellerByUserID(userIDFromJWT)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Phone != "" {
		seller.Phone = req.Phone
	}
	if req.StoreName != "" {
		seller.StoreName = req.StoreName
		// Generate new slug if store name changed
		seller.StoreSlug = util.GenerateSlug(req.StoreName)
	}
	if req.StoreDesc != "" {
		seller.StoreDesc = req.StoreDesc
	}
	if req.BusinessType != "" {
		seller.BusinessType = req.BusinessType
	}

	if req.Address != "" {
		seller.Address = req.Address
	}
	if req.City != "" {
		seller.City = req.City
	}
	if req.State != "" {
		seller.State = req.State
	}
	if req.Country != "" {
		seller.Country = req.Country
	}
	if req.PostalCode != "" {
		seller.PostalCode = req.PostalCode
	}

	// Save updates
	if err := s.sellerRepo.UpdateSellerProfile(seller); err != nil {
		return nil, fmt.Errorf("failed to update profile: %v", err)
	}

	// Return updated profile
	return s.GetSellerProfile(userIDFromJWT)
}
