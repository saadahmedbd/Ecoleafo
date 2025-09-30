package sellerprofileservice

import (
	"fmt"

	sellerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/SellerProfile"
)

func (s *sellerService) GetPublicSellerProfile(sellerID uint) (*sellerprofile.SellerProfileResponse, error) {
	// Get seller by primary key ID
	seller, err := s.sellerRepo.GetSellerByID(sellerID)
	if err != nil {
		return nil, err
	}

	// Only return if seller is active and approved
	if !seller.IsActive || !seller.IsApproved {
		return nil, fmt.Errorf("seller profile not available")
	}

	// Get seller with relations
	sellerWithRelations, err := s.sellerRepo.GetSellerWithRelations(seller.ID)
	if err != nil {
		return nil, err
	}

	return s.mapToSellerProfileResponse(sellerWithRelations), nil
}
