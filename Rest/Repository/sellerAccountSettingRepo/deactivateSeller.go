package selleraccountsettingrepo

import models "github.com/saadahmedbd/Treestore/Models"

// DeactivateSeller deactivates seller account
func (r *SellerAccountSettingRepository) DeactivateSeller(sellerID uint) error {
	return r.db.Model(&models.User{}).
		Where("id = ?", sellerID).
		Update("is_active", false).Error
}
