package selleraccountsettingrepo

import models "github.com/saadahmedbd/Treestore/Models"

// Toggle2FA enables or disables 2FA for seller
func (r *SellerAccountSettingRepository) Toggle2FA(userID uint, enabled bool) error {
	return r.db.Model(&models.RegUser{}).
		Where("id = ?", userID).
		Update("two_factor_enabled", enabled).Error
}
