package buyerservice

import (
	"fmt"

	util "github.com/saadahmedbd/Treestore/Util"
)

// DeleteBuyerProfilePicture deletes a profile picture from Cloudinary and DB
func (s *buyerService) DeleteBuyerProfilePicture(userID uint, imageURL string) error {
	// Step 1: Delete from Cloudinary
	if err := util.DeleteImageFromCloudinary(imageURL); err != nil {
		return fmt.Errorf("failed to delete from Cloudinary: %w", err)
	}

	// Step 2: Clear profile photo in DB
	if err := s.buyerRepo.ClearProfilePhoto(userID); err != nil {
		return fmt.Errorf("failed to clear profile image in DB: %w", err)
	}

	return nil
}
