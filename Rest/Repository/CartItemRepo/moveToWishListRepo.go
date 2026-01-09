package cartitemrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *cartRepository) MoveToWishlist(buyerID, productID uint) error {
	tx := r.db.Begin()

	// Check if already in wishlist
	var existing models.Wishlist
	err := tx.Where("buyer_id = ? AND product_id = ?", buyerID, productID).First(&existing).Error
	if err == nil {
		// Already in wishlist, just remove from cart
		if err := tx.Where("buyer_id = ? AND product_id = ?", buyerID, productID).
			Delete(&models.CartItem{}).Error; err != nil {
			tx.Rollback()
			return err
		}
		return tx.Commit().Error
	}

	// Add to wishlist
	wishlist := &models.Wishlist{
		BuyerID:   buyerID,
		ProductID: productID,
		AddedFrom: "cart",
	}

	if err := tx.Create(wishlist).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Remove from cart
	if err := tx.Where("buyer_id = ? AND product_id = ?", buyerID, productID).
		Delete(&models.CartItem{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
