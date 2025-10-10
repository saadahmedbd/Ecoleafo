package cartitemrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *cartRepository) ClearCart(buyerID uint) error {
	return r.db.Where("buyer_id = ? AND is_saved_for_later = ?", buyerID, false).
		Delete(&models.CartItem{}).Error
}
