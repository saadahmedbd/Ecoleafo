package selleraccountsettingrepo

import (
	"errors"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

// GetRegUserByID retrieves RegUser by ID
func (r *SellerAccountSettingRepository) GetRegUserByID(userID uint) (*models.RegUser, error) {
	var regUser models.RegUser
	err := r.db.First(&regUser, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &regUser, nil
}
