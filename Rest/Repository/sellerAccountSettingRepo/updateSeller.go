package selleraccountsettingrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
)

// UpdateSeller updates seller information
func (r *SellerAccountSettingRepository) UpdateSeller(sellerID uint, updates map[string]interface{}) error {
	updates["updated_at"] = time.Now()
	return r.db.Model(&models.User{}).Where("id = ?", sellerID).Updates(updates).Error
}
