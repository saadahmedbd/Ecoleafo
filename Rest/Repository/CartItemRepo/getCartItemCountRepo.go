package cartitemrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *cartRepository) GetCartItemCount(buyerID uint) (int, error) {
	var count int64
	err := r.db.Model(&models.CartItem{}).
		Where("buyer_id = ? AND is_saved_for_later = ?", buyerID, false).
		Count(&count).Error
	return int(count), err
}
func (r *cartRepository) GetWishlistCount(buyerID uint) (int, error) {
	var count int64
	err := r.db.Model(&models.Wishlist{}).
		Where("buyer_id = ?", buyerID).
		Count(&count).Error
	return int(count), err
}
