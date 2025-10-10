package cartitemrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *cartRepository) RemoveFromWishlist(buyerID, productID uint) error {
	return r.db.Where("buyer_id = ? AND product_id = ?", buyerID, productID).
		Delete(&models.Wishlist{}).Error
}
