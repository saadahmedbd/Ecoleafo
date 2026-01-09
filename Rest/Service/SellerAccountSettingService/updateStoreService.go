package selleraccountsettingservice

import (
	"errors"

	selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"
	util "github.com/saadahmedbd/Treestore/Util"
)

// UpdateStore updates store information
func (s *SellerAccountSettingService) UpdateStore(sellerID uint, req *selleraccountsetting.UpdateStoreRequest) error {
	// Validate request
	if err := util.ValidateStruct(req); err != nil {
		return err
	}

	// Generate store slug if store name changed
	storeSlug := util.GenerateSlug(req.StoreName)

	// Prepare updates
	updates := map[string]interface{}{
		"store_name":     req.StoreName,
		"store_slug":     storeSlug,
		"store_desc":     req.StoreDescription,
		"business_email": req.BusinessEmail,
		"phone":          req.Phone,
		"address":        req.Address,
		"city":           req.City,
		"state":          req.State,
		"country":        req.Country,
		"postal_code":    req.PostalCode,
	}

	// Update seller
	if err := s.repo.UpdateSeller(sellerID, updates); err != nil {
		return errors.New("failed to update store information")
	}

	return nil
}
