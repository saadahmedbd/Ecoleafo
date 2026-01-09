package selleraccountsettingrepo

import models "github.com/saadahmedbd/Treestore/Models"

// GetLoginActivities retrieves recent login activities
func (r *SellerAccountSettingRepository) GetLoginActivities(sellerID uint, limit int) ([]models.SellerLoginActivity, error) {
	var activities []models.SellerLoginActivity
	err := r.db.Where("seller_id = ?", sellerID).
		Order("created_at DESC").
		Limit(limit).
		Find(&activities).Error
	return activities, err
}
