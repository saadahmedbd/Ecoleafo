package buyerservice

import (
	"errors"

	buyerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/BuyerProfile"
)

func (s *buyerService) UpdateProfilePhoto(buyerID uint, photoURL, publicID string) (*buyerprofile.UpdateProfilePhotoResponse, error) {
	buyer, err := s.buyerRepo.GetBuyerByID(buyerID)
	if err != nil {
		return nil, err
	}
	// Update profile photo
	updates := map[string]interface{}{
		"profile_picture_url":       photoURL,
		"profile_picture_public_id": publicID,
	}

	if err := s.buyerRepo.UpdateRegUserForImage(buyer.UserId, updates); err != nil {
		return nil, errors.New("failed to update profile photo")
	}

	return &buyerprofile.UpdateProfilePhotoResponse{
		PhotoURL: photoURL,
		PublicID: publicID,
		Message:  "Profile photo updated successfully",
	}, nil
}
