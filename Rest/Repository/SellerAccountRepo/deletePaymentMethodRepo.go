package selleraccountrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *sellerRegistrationRepository) DeletePaymentMethod(paymentMethodID, sellerID uint) error {
	return r.db.Where("id = ? AND seller_id = ?", paymentMethodID, sellerID).
		Delete(&models.SellerPaymentMethod{}).Error
}
