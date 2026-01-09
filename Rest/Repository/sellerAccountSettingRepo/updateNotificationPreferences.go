package selleraccountsettingrepo

import (
	"errors"
	"fmt"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

// UpdateNotificationPreferences updates notification settings
func (r *SellerAccountSettingRepository) UpdateNotificationPreferences(jwtUserID uint, updates map[string]interface{}) error {
	// Step 1: Get the actual seller ID from user_id
	var seller models.User
	if err := r.db.Where("user_id = ?", jwtUserID).First(&seller).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("seller not found for user_id %d", jwtUserID)
		}
		return err
	}

	sellerID := seller.ID
	updates["updated_at"] = time.Now()

	// Step 2: Check if preferences exist
	var count int64
	r.db.Model(&models.SellerNotificationPreference{}).Where("seller_id = ?", sellerID).Count(&count)

	if count == 0 {
		// Create default preferences if not exist
		prefs := models.SellerNotificationPreference{
			SellerID:     sellerID,
			OrderEmail:   true,
			OrderPush:    true,
			MessageEmail: true,
			MessagePush:  true,
		}
		return r.db.Create(&prefs).Error
	}

	// Step 3: Update existing preferences
	return r.db.Model(&models.SellerNotificationPreference{}).
		Where("seller_id = ?", sellerID).
		Updates(updates).Error
}
