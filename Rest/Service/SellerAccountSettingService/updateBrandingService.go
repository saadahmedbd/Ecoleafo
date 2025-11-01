package selleraccountsettingservice

import (
	"errors"

	selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"
)

// UpdateBranding updates store logo and banner
func (s *SellerAccountSettingService) UpdateBranding(sellerID uint, logoURL, bannerURL string) (*selleraccountsetting.UpdateBrandingResponse, error) {
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

	// Update seller
	if err := s.repo.UpdateSeller(sellerID, updates); err != nil {
		return nil, errors.New("failed to update branding")
	}

	return &selleraccountsetting.UpdateBrandingResponse{
		LogoURL:   logoURL,
		BannerURL: bannerURL,
		Message:   "Store branding updated successfully",
	}, nil
}
