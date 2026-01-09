package guestcartservice

import (
	buyeraccount "github.com/saadahmedbd/Treestore/Rest/DTO/BuyerAccount"
)

func (s *guestcartservice) CheckProfileCompletion(userID uint) (*buyeraccount.ProfileCompletionStatus, error) {
	buyer, err := s.buyerRepo.GetBuyerByUserID(userID)
	if err != nil {
		// If buyer doesn't exist, profile is incomplete
		if err.Error() == "buyer not found" {
			return &buyeraccount.ProfileCompletionStatus{
				IsProfileComplete: false,
				HasPhone:          false,
				HasAddress:        false,
				MissingFields:     []string{"Phone", "Address"},
				NextStep:          "complete_profile",
			}, nil
		}
		return nil, err
	}

	isComplete, missingFields, err := s.profileRepo.HasCompletedProfile(userID)
	if err != nil {
		return nil, err
	}

	status := &buyeraccount.ProfileCompletionStatus{
		IsProfileComplete: isComplete,
		HasPhone:          buyer.Phone != "",
		MissingFields:     missingFields,
	}

	// Check if has address
	addresses, _ := s.profileRepo.GetBuyerAddresses(buyer.ID)
	status.HasAddress = len(addresses) > 0

	// Determine next step
	if !status.HasPhone {
		status.NextStep = "complete_profile"
	} else if !status.HasAddress {
		status.NextStep = "add_address"
	} else {
		status.NextStep = "ready_to_checkout"
	}

	return status, nil
}
