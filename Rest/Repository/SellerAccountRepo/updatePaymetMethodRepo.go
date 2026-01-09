package selleraccountrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *sellerRegistrationRepository) UpdatePaymentMethod(paymentMethod *models.SellerPaymentMethod) error {
	if paymentMethod.IsDefault {
		r.db.Model(&models.SellerPaymentMethod{}).
			Where("seller_id = ? AND id != ?", paymentMethod.SellerID, paymentMethod.ID).
			Update("is_default", false)
	}
	return r.db.Save(paymentMethod).Error
}
