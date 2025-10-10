package database

import "gorm.io/gorm"

// Composite unique index to prevent duplicates
type CartItemIndex struct {
}

func (CartItemIndex) Migrate(db *gorm.DB) error {
	// Add unique index on buyer_id and product_id (excluding saved for later)
	return db.Exec(`
		CREATE UNIQUE INDEX IF NOT EXISTS idx_cart_items_buyer_product 
		ON cart_items(buyer_id, product_id) 
		WHERE is_saved_for_later = false
	`).Error
}
