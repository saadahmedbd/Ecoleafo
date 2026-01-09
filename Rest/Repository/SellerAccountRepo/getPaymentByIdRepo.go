package selleraccountrepo

import (
	"errors"
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *sellerRegistrationRepository) GetPaymentMethodByID(paymentMethodID, sellerID uint) (*models.SellerPaymentMethod, error) {
	var paymentMethod models.SellerPaymentMethod
	err := r.db.Where("id = ? AND seller_id = ?", paymentMethodID, sellerID).First(&paymentMethod).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("payment method not found")
		}
		return nil, err
	}
	return &paymentMethod, nil

}
