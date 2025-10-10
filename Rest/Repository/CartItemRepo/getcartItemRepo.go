package cartitemrepo

import (
	"errors"
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *cartRepository) GetCartItem(buyerID, productID uint) (*models.CartItem, error) {
	var item models.CartItem
	err := r.db.Where("buyer_id = ? AND product_id = ?", buyerID, productID).First(&item).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("item not found in cart")
		}
		return nil, err
	}
	return &item, nil
}
