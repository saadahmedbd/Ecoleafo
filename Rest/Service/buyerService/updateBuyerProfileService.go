package buyerservice

import (
	"fmt"
	"time"

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
	if req.Gender != "" {
		buyer.RegUser.Gender = req.Gender
	}
	if req.DateOfBirth != "" {
		parsedDate, err := time.Parse("2006-01-02T15:04:05.000Z", req.DateOfBirth)
		if err != nil {
			return nil, fmt.Errorf("invalid date format: %v", err)
		}
		buyer.RegUser.DateOfBirth = &parsedDate
	}

	if err := s.buyerRepo.UpdateBuyerProfile(buyer); err != nil {
		return nil, fmt.Errorf("failed to update profile: %v", err)
	}

	return s.GetBuyerProfile(userIDFromJWT)
}
