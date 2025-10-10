package cartitemrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *cartRepository) UpdateQuantity(buyerID, productID uint, quantity int) error {
	return r.db.Model(&models.CartItem{}).
		Where("buyer_id = ? AND product_id = ?", buyerID, productID).
		Update("quantity", quantity).Error
}
