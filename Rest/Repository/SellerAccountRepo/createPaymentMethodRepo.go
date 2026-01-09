package selleraccountrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *sellerRegistrationRepository) CreatePaymentMethod(paymentMethod *models.SellerPaymentMethod) error {
	if paymentMethod.IsDefault {
		r.db.Model(&models.SellerPaymentMethod{}).
			Where("seller_id = ?", paymentMethod.SellerID).
			Update("is_default", false)

	}
	return r.db.Create(paymentMethod).Error
}
