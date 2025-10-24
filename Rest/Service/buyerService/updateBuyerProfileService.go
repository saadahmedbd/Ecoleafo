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
	if req.FirstName != "" {
		buyer.RegUser.FirstName = req.FirstName
	}
	if req.LastName != "" {
		buyer.RegUser.LastName = req.LastName
	}
	if req.Email != "" {
		buyer.RegUser.Email = req.Email
	}
	if req.ProfilePicture != "" {
		buyer.ProfilePicture = req.ProfilePicture
	}

	if err := s.buyerRepo.UpdateBuyerProfile(buyer); err != nil {
		return nil, fmt.Errorf("failed to update profile: %v", err)
	}

	return s.GetBuyerProfile(userIDFromJWT)
}
