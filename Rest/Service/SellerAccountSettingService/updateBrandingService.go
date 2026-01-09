package selleraccountsettingservice

import (
	"errors"

	selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"
)

// UpdateBranding updates store logo and banner
func (s *SellerAccountSettingService) UpdateBranding(userID uint, logoURL, bannerURL string) (*selleraccountsetting.UpdateBrandingResponse, error) {
	// Get seller by RegUser ID (from JWT)
	seller, err := s.repo.GetSellerByUserID(userID)
	if err != nil {
		return nil, err
	}

	updates := make(map[string]interface{})

	if logoURL != "" {
		updates["store_logo"] = logoURL
	}

	if bannerURL != "" {
		updates["store_banner"] = bannerURL
	}

	if len(updates) == 0 {
		return nil, errors.New("no branding files provided")
	}

	// Update seller using actual seller ID
	if err := s.repo.UpdateSeller(seller.ID, updates); err != nil {
		return nil, errors.New("failed to update branding")
	}

	return &selleraccountsetting.UpdateBrandingResponse{
		LogoURL:   logoURL,
		BannerURL: bannerURL,
		Message:   "Store branding updated successfully",
	}, nil
}
