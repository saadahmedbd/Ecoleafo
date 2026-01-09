package selleraccountsettingrepo

import models "github.com/saadahmedbd/Treestore/Models"

// UpdatePassword updates seller password
func (r *SellerAccountSettingRepository) UpdatePassword(userID uint, hashedPassword string) error {
	return r.db.Model(&models.RegUser{}).
		Where("id = ?", userID).
		Update("password", hashedPassword).Error
}
