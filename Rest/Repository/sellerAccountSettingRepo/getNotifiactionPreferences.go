package selleraccountsettingrepo

import (
	"errors"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

// GetNotificationPreferences retrieves notification settings
func (r *SellerAccountSettingRepository) GetNotificationPreferences(sellerID uint) (*models.SellerNotificationPreference, error) {
	var prefs models.SellerNotificationPreference
	err := r.db.Where("seller_id = ?", sellerID).First(&prefs).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Create default preferences if not found
			prefs = models.SellerNotificationPreference{
				SellerID:     sellerID,
				OrderEmail:   true,
				OrderPush:    true,
				MessageEmail: true,
				MessagePush:  true,
			}
			if err := r.db.Create(&prefs).Error; err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return &prefs, nil
}
