package selleraccountsettingrepo

import (
	"errors"
	"fmt"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

// UpdateSellerPolicies updates seller policies
func (r *SellerAccountSettingRepository) UpdateSellerPolicies(jwtUserID uint, updates map[string]interface{}) error {
	updates["updated_at"] = time.Now()

	// Map JWT user_id → actual seller ID in users table
	var seller models.User
	if err := r.db.Where("user_id = ?", jwtUserID).First(&seller).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("seller not found for user_id %d", jwtUserID)
		}
		return err
	}
	sellerID := seller.ID

	// Check if policies exist
	var count int64
	r.db.Model(&models.SellerPolicy{}).Where("seller_id = ?", sellerID).Count(&count)

	if count == 0 {
		// Create new policies
		policy := models.SellerPolicy{
			SellerID:       sellerID,
			ReturnPolicy:   fmt.Sprintf("%v", updates["return_policy"]),
			ShippingPolicy: fmt.Sprintf("%v", updates["shipping_policy"]),
			FAQ:            fmt.Sprintf("%v", updates["faq"]),
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}
		return r.db.Create(&policy).Error
	}

	// Update existing policies
	return r.db.Model(&models.SellerPolicy{}).
		Where("seller_id = ?", sellerID).
		Updates(updates).Error
}
