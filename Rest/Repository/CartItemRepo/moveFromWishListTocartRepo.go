package cartitemrepo

import (
	"errors"
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *cartRepository) MoveFromWishlistToCart(buyerID, productID uint, quantity int) error {
	// Check if item exists in wishlist
	var wishlistItem models.Wishlist
	if err := r.db.Where("buyer_id = ? AND product_id = ?", buyerID, productID).First(&wishlistItem).Error; err != nil {
		return fmt.Errorf("item not in wishlist: %w", err)
	}

	// Get product price
	var product models.Product
	if err := r.db.Where("id = ?", productID).First(&product).Error; err != nil {
		return fmt.Errorf("product not found: %w", err)
	}

	// Check if already in cart
	var existing models.CartItem
	err := r.db.Where("buyer_id = ? AND product_id = ? AND is_saved_for_later = ?", buyerID, productID, false).First(&existing).Error
	if err == nil {
		// Update quantity
		existing.Quantity += quantity
		if err := r.db.Save(&existing).Error; err != nil {
			return fmt.Errorf("failed to update cart: %w", err)
		}
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		// Create new cart item
		cartItem := &models.CartItem{
			BuyerID:   buyerID,
			ProductID: productID,
			Quantity:  quantity,
			Price:     product.Price,
			AddedFrom: "wishlist",
		}
		if err := r.db.Create(cartItem).Error; err != nil {
			return fmt.Errorf("failed to create cart item: %w", err)
		}
	} else {
		return fmt.Errorf("failed to check cart: %w", err)
	}

	// Remove from wishlist
	if err := r.db.Where("buyer_id = ? AND product_id = ?", buyerID, productID).
		Delete(&models.Wishlist{}).Error; err != nil {
		return fmt.Errorf("failed to remove from wishlist: %w", err)
	}

	return nil
}
