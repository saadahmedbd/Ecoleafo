package selleraccountsettingrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
)

// UpdateNotificationPreferences updates notification settings
func (r *SellerAccountSettingRepository) UpdateNotificationPreferences(sellerID uint, updates map[string]interface{}) error {
	updates["updated_at"] = time.Now()

	// Check if preferences exist
	var count int64
	r.db.Model(&models.SellerNotificationPreference{}).Where("seller_id = ?", sellerID).Count(&count)

	if count == 0 {
		// Create new preferences
		prefs := models.SellerNotificationPreference{SellerID: sellerID}
		return r.db.Create(&prefs).Error
	}

	return r.db.Model(&models.SellerNotificationPreference{}).
		Where("seller_id = ?", sellerID).
		Updates(updates).Error
}
