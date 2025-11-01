package selleraccountsettingrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
)

// UpdateSellerPolicies updates seller policies
func (r *SellerAccountSettingRepository) UpdateSellerPolicies(sellerID uint, updates map[string]interface{}) error {
	updates["updated_at"] = time.Now()

	// Check if policies exist
	var count int64
	r.db.Model(&models.SellerPolicy{}).Where("seller_id = ?", sellerID).Count(&count)

	if count == 0 {
		// Create new policies
		policy := models.SellerPolicy{SellerID: sellerID}
		return r.db.Create(&policy).Error
	}

	return r.db.Model(&models.SellerPolicy{}).
		Where("seller_id = ?", sellerID).
		Updates(updates).Error
}
