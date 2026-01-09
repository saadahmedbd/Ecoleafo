package selleraccountsettingrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
)

func (r *SellerAccountSettingRepository) UpdateCommission(sellerID uint, commission float64) error {
	return r.db.Model(&models.User{}).
		Where("id = ?", sellerID).
		Update("commission", commission).Error
}
