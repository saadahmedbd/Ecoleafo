package buyerservice

import (
	"errors"

	buyerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/BuyerProfile"
)

func (s *buyerService) UpdateProfilePhoto(buyerID uint, photoURL string) (*buyerprofile.UpdateProfilePhotoResponse, error) {
	buyer, err := s.buyerRepo.GetBuyerByID(buyerID)
	if err != nil {
		return nil, err
	}
	// Update profile photo
	updates := map[string]interface{}{
		"profile_photo": photoURL,
	}

	if err := s.buyerRepo.UpdateRegUserForImage(buyer.UserId, updates); err != nil {
		return nil, errors.New("failed to update profile photo")
	}

	return &buyerprofile.UpdateProfilePhotoResponse{
		PhotoURL: photoURL,
		Message:  "Profile photo updated successfully",
	}, nil
}
