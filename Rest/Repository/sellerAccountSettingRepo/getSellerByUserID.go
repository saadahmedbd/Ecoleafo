package selleraccountsettingrepo

import (
	"errors"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

// GetSellerByUserID retrieves seller by RegUser ID
func (r *SellerAccountSettingRepository) GetSellerByUserID(userID uint) (*models.User, error) {
	var seller models.User
	err := r.db.Preload("RegUser").Where("user_id = ?", userID).First(&seller).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("seller not found")
		}
		return nil, err
	}
	return &seller, nil
}
