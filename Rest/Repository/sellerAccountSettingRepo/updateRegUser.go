package selleraccountsettingrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
)

// UpdateRegUser updates RegUser information
func (r *SellerAccountSettingRepository) UpdateRegUser(userID uint, updates map[string]interface{}) error {
	updates["updated_at"] = time.Now()
	return r.db.Model(&models.RegUser{}).Where("id = ?", userID).Updates(updates).Error
}
