package selleraccountrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *sellerRegistrationRepository) GetPaymentMethods(sellerID uint) ([]models.SellerPaymentMethod, error) {
	var paymentMethod []models.SellerPaymentMethod
	err := r.db.Where("seller_id = ?", sellerID).
		Order("is_default DESC, created_at DESC").
		Find(&paymentMethod).Error
	if err != nil {
		return nil, err
	}
	return paymentMethod, nil
}
