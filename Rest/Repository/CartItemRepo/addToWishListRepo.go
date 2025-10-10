package cartitemrepo

import (
	"errors"
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *cartRepository) AddToWishlist(wishlist *models.Wishlist) error {
	// Check if already in wishlist
	var existing models.Wishlist
	err := r.db.Where("buyer_id = ? AND product_id = ?", wishlist.BuyerID, wishlist.ProductID).
		First(&existing).Error

	if err == nil {
		return fmt.Errorf("product already in wishlist")
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return r.db.Create(wishlist).Error
	}

	return err
}
