package selleraccountsettingrepo

import (
	"errors"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

// GetSellerByID retrieves seller by ID with related data
func (r *SellerAccountSettingRepository) GetSellerByID(sellerID uint) (*models.User, error) {
	var seller models.User
	err := r.db.Preload("RegUser").Where("user_id = ?", sellerID).First(&seller).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("seller not found")
		}
		return nil, err
	}
	return &seller, nil
}
