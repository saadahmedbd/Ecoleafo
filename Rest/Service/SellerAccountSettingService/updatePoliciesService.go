package selleraccountsettingservice

import (
	"errors"

	selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"
	util "github.com/saadahmedbd/Treestore/Util"
)

// UpdatePolicies updates store policies
func (s *SellerAccountSettingService) UpdatePolicies(sellerID uint, req *selleraccountsetting.UpdatePoliciesRequest) error {
	// Validate request
	if err := util.ValidateStruct(req); err != nil {
		return err
	}

	// Prepare updates
	updates := map[string]interface{}{}

	if req.ReturnPolicy != "" {
		updates["return_policy"] = req.ReturnPolicy
	}

	if req.ShippingPolicy != "" {
		updates["shipping_policy"] = req.ShippingPolicy
	}

	if req.FAQ != "" {
		updates["faq"] = req.FAQ
	}

	if len(updates) == 0 {
		return errors.New("no policies to update")
	}

	// Update policies
	if err := s.repo.UpdateSellerPolicies(sellerID, updates); err != nil {
		return errors.New("failed to update policies")
	}

	return nil
}
