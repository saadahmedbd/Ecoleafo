package cartitemrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *cartRepository) SaveForLater(buyerID, productID uint) error {
	result := r.db.Model(&models.CartItem{}).
		Where("buyer_id = ? AND product_id = ? AND is_saved_for_later = ?", buyerID, productID, false).
		Updates(map[string]interface{}{"is_saved_for_later": true, "updated_at": time.Now()})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
