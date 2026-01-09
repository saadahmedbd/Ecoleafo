package selleraccountsettingrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
)

// DeleteRegUser soft deletes RegUser account
func (r *SellerAccountSettingRepository) DeleteRegUser(userID uint) error {
	now := time.Now()
	return r.db.Model(&models.RegUser{}).
		Where("id = ?", userID).
		Update("deleted_at", &now).Error
}
