package models

import "time"

type Wishlist struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	BuyerID   uint      `json:"buyer_id" gorm:"not null;uniqueIndex:idx_buyer_product"`
	ProductID uint      `json:"product_id" gorm:"not null;uniqueIndex:idx_buyer_product"`
	Quantity  uint      `json:"quantity" gorm:"default:1"` // only if it's a cart
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Buyer   Buyer   `json:"buyer" gorm:"foreignKey:BuyerID"`
	Product Product `json:"product" gorm:"foreignKey:ProductID"`
}
