package selleraccountrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *sellerRegistrationRepository) SetDefaultPaymentMethod(sellerID, paymentMethodID uint) error {
	tx := r.db.Begin()

	if err := tx.Model(&models.SellerPaymentMethod{}).
		Where("seller_id = ?", sellerID).
		Update("is_default", false).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&models.SellerPaymentMethod{}).
		Where("id = ? AND seller_id = ?", paymentMethodID, sellerID).
		Update("is_default", true).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
