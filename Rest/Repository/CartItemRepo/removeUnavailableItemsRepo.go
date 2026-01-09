package cartitemrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *cartRepository) RemoveUnavailableItems(buyerID uint) error {
	// Remove items where product is inactive or out of stock
	return r.db.Where(`buyer_id = ? AND product_id IN (
		SELECT id FROM products WHERE is_active = false OR quantity = 0
	)`, buyerID).Delete(&models.CartItem{}).Error
}
