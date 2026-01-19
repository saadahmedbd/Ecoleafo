package selleraccountrepo

import (
	"errors"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *sellerRegistrationRepository) GetSellerByEmail(email string) (*models.User, error) {
	var seller models.User
	err := r.db.Where("business_email = ?", email).First(&seller).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &seller, nil
}
