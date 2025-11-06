package selleraccountrepo

import (
	"errors"
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *sellerRegistrationRepository) GetSellerByRegUserId(regUserID uint) (*models.User, error) {
	var seller models.User
	err := r.db.Where("user_id", regUserID).First(&seller).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("seller not found")
		}
		return nil, err
	}
	return &seller, nil
}
