package selleraccountsettingrepo

import models "github.com/saadahmedbd/Treestore/Models"

// SetCurrentLoginActivity marks a login as current and unmarks others
func (r *SellerAccountSettingRepository) SetCurrentLoginActivity(sellerID uint, activityID uint) error {
	// Start transaction
	tx := r.db.Begin()

	// Unmark all current sessions
	if err := tx.Model(&models.SellerLoginActivity{}).
		Where("seller_id = ?", sellerID).
		Update("is_current", false).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Mark new current session
	if err := tx.Model(&models.SellerLoginActivity{}).
		Where("id = ?", activityID).
		Update("is_current", true).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
