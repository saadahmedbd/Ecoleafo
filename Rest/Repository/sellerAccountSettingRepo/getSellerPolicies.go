package selleraccountsettingrepo

import (
	"errors"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

// GetSellerPolicies retrieves seller policies
func (r *SellerAccountSettingRepository) GetSellerPolicies(sellerID uint) (*models.SellerPolicy, error) {
	var policy models.SellerPolicy
	err := r.db.Where("user_id = ?", sellerID).First(&policy).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Create default policies if not found
			policy = models.SellerPolicy{SellerID: sellerID}
			if err := r.db.Create(&policy).Error; err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return &policy, nil
}
