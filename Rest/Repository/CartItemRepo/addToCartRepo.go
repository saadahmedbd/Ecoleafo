package cartitemrepo

import (
	"errors"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *cartRepository) AddToCart(item *models.CartItem) error {
	// Check if item already exists for this buyer
	var existing models.CartItem

	err := r.db.
		Where("buyer_id = ? AND product_id = ? AND is_saved_for_later = FALSE",
			item.BuyerID, item.ProductID).
		First(&existing).Error

	if err == nil {
		// Item exists -> update quantity
		existing.Quantity += item.Quantity
		existing.UpdatedAt = time.Now()
		return r.db.Save(&existing).Error
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// New item -> insert
		item.CreatedAt = time.Now()
		item.UpdatedAt = time.Now()
		return r.db.Create(item).Error
	}

	// Other DB error
	return err
}
