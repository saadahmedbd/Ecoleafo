package selleraccountsettingrepo

import models "github.com/saadahmedbd/Treestore/Models"

// Get2FAStatus checks if 2FA is enabled for seller
func (r *SellerAccountSettingRepository) Get2FAStatus(userID uint) (bool, error) {
	var regUser models.RegUser
	err := r.db.Select("two_factor_enabled").First(&regUser, userID).Error
	if err != nil {
		return false, err
	}
	return regUser.TwoFactorEnabled, nil
}
