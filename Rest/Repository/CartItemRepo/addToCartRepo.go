package cartitemrepo

import (
	"errors"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *cartRepository) AddToCart(item *models.CartItem) error {
	//chck if item alreday exists

	var existing models.CartItem
	err := r.db.Where("buyer_id = ? AND product_id = ? AND is_saved_for_later = ?",
		item.BuyerID, item.ProductID, false).First(&existing).Error
	if err == nil {
		// Item exists, update quantity
		existing.Quantity += item.Quantity
		existing.UpdatedAt = time.Now()
		return r.db.Save(&existing).Error
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//create new item
		return r.db.Create(item).Error
	}
	return err

}
