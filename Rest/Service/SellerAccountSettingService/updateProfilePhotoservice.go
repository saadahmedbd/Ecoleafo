package selleraccountsettingservice

import (
	"errors"

	selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"
)

// UpdateProfilePhoto updates seller profile photo
func (s *SellerAccountSettingService) UpdateProfilePhoto(sellerID uint, photoURL string) (*selleraccountsetting.UpdateProfilePhotoResponse, error) {
	// Get seller to find RegUser ID
	seller, err := s.repo.GetSellerByID(sellerID)
	if err != nil {
		return nil, err
	}

	// Update profile photo
	updates := map[string]interface{}{
		"profile_photo": photoURL,
	}

	if err := s.repo.UpdateRegUser(seller.UserId, updates); err != nil {
		return nil, errors.New("failed to update profile photo")
	}

	return &selleraccountsetting.UpdateProfilePhotoResponse{
		PhotoURL: photoURL,
		Message:  "Profile photo updated successfully",
	}, nil
}
