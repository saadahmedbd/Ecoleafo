package cartitemrepo

import (
	"errors"

	models "github.com/saadahmedbd/Treestore/Models"
)

func (r *cartRepository) RemoveFromWishlist(buyerID, productID uint) error {
	result := r.db.Where("buyer_id = ? AND product_id = ?", buyerID, productID).
		Delete(&models.Wishlist{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("item not found in wishlist")
	}

	return nil
}
