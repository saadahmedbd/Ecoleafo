package buyerservice

import (
	"fmt"

	buyerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/BuyerProfile"
)

func (s *buyerService) UpdateBuyerProfile(userIDFromJWT uint, req buyerprofile.UpdateBuyerProfileRequest) (*buyerprofile.BuyerProfileResponse, error) {
	buyer, err := s.buyerRepo.GetBuyerByUserID(userIDFromJWT)
	if err != nil {
		return nil, err
	}

	if req.Phone != "" {
		buyer.Phone = req.Phone
	}
	if req.DefaultAddress != "" {
		buyer.DefaultAddress = req.DefaultAddress
	}

	if err := s.buyerRepo.UpdateBuyerProfile(buyer); err != nil {
		return nil, fmt.Errorf("failed to update profile: %v", err)
	}

	return s.GetBuyerProfile(userIDFromJWT)
}
