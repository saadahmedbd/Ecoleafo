package cartitemrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *cartRepository) GetCartTotal(buyerID uint) (float64, error) {
	var total float64

	// Use COALESCE to ensure NULL becomes 0
	err := r.db.Model(&models.CartItem{}).
		Select("COALESCE(SUM(price * quantity), 0)").
		Where("buyer_id = ? AND is_saved_for_later = ?", buyerID, false).
		Scan(&total).Error

	if err != nil {
		return 0, err
	}

	return total, nil
}
