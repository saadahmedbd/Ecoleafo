package selleraccountsettingrepo

import (
	"errors"
	"fmt"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

// GetNotificationPreferences retrieves notification settings
func (r *SellerAccountSettingRepository) GetNotificationPreferences(jwtUserID uint) (*models.SellerNotificationPreference, error) {
	// Step 1: Find the actual seller ID from users table
	var seller models.User
	if err := r.db.Where("user_id = ?", jwtUserID).First(&seller).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("seller not found for user_id %d", jwtUserID)
		}
		return nil, err
	}
	sellerID := seller.ID

	// Step 2: Get notification preferences
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
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
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
