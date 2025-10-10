package cartitemrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *cartRepository) MoveToCart(buyerID, productID uint) error {
	result := r.db.Model(&models.CartItem{}).
		Where("buyer_id = ? AND product_id = ? AND is_saved_for_later = ?", buyerID, productID, true).
		Updates(map[string]interface{}{"is_saved_for_later": false, "updated_at": time.Now()})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
