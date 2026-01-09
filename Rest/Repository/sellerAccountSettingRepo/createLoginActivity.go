package selleraccountsettingrepo

import models "github.com/saadahmedbd/Treestore/Models"

// CreateLoginActivity records login activity
func (r *SellerAccountSettingRepository) CreateLoginActivity(activity *models.SellerLoginActivity) error {
	return r.db.Create(activity).Error
}
