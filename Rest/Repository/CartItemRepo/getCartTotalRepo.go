package cartitemrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *cartRepository) GetCartTotal(buyerID uint) (float64, error) {
	var total float64
	err := r.db.Model(&models.CartItem{}).
		Select("SUM(price * quantity)").
		Where("buyer_id = ? AND is_saved_for_later = ?", buyerID, false).
		Scan(&total).Error
	return total, err
}
