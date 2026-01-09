package selleraccountsettingrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
)

// DeleteSeller soft deletes seller account
func (r *SellerAccountSettingRepository) DeleteSeller(sellerID uint) error {
	now := time.Now()
	return r.db.Model(&models.User{}).
		Where("id = ?", sellerID).
		Update("deleted_at", &now).Error
}
