package cartitemrepo

import (
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *cartRepository) GetSelectedCartItems(buyerID uint) ([]models.CartItem, error) {
	var items []models.CartItem
	err := r.db.Preload("Product").
		Preload("Product.Images").
		Preload("Product.Seller", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, store_name")
		}).
		Where("buyer_id = ? AND is_selected = ? AND is_saved_for_later = ?", buyerID, true, false).
		Order("created_at DESC").
		Find(&items).Error
	
	// Debug log
	if len(items) == 0 {
		// Check if any items exist
		var count int64
		r.db.Model(&models.CartItem{}).Where("buyer_id = ?", buyerID).Count(&count)
		if count > 0 {
			// Items exist but none selected
			var selectedCount int64
			r.db.Model(&models.CartItem{}).Where("buyer_id = ? AND is_selected = ?", buyerID, true).Count(&selectedCount)
			fmt.Printf("Cart items for buyer %d: total=%d, selected=%d\n", buyerID, count, selectedCount)
		}
	}
	
	return items, err
}

func (r *cartRepository) UpdateCartItemSelection(buyerID uint, productID uint, isSelected bool) error {
	return r.db.Table("cart_items").
		Where("buyer_id = ? AND product_id = ?", buyerID, productID).
		Update("is_selected", isSelected).Error
}

func (r *cartRepository) UpdateCartItemSelectionByID(buyerID uint, cartItemID uint, isSelected bool) error {
	return r.db.Table("cart_items").
		Where("buyer_id = ? AND id = ?", buyerID, cartItemID).
		Update("is_selected", isSelected).Error
}

func (r *cartRepository) UpdateAllCartItemsSelection(buyerID uint, isSelected bool) error {
	return r.db.Table("cart_items").
		Where("buyer_id = ?", buyerID).
		Update("is_selected", isSelected).Error
}
