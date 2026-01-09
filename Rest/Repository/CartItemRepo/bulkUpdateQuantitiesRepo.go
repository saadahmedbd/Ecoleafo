package cartitemrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *cartRepository) BulkUpdateQuantities(buyerID uint, updates map[uint]int) error {
	tx := r.db.Begin()

	for productID, quantity := range updates {
		if quantity <= 0 {
			if err := tx.Where("buyer_id = ? AND product_id = ?", buyerID, productID).
				Delete(&models.CartItem{}).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			if err := tx.Model(&models.CartItem{}).
				Where("buyer_id = ? AND product_id = ?", buyerID, productID).
				Update("quantity", quantity).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}
