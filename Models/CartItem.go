package models

import "time"

// CartItem represents items in buyer's shopping cart - MVP
type CartItem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	BuyerID   uint      `json:"buyer_id" gorm:"not null"`
	ProductID uint      `json:"product_id" gorm:"not null"`
	Quantity  int       `json:"quantity" gorm:"not null;check:quantity > 0"`
	Price     float64   `json:"price" gorm:"type:decimal(10,2);not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relationships
	Buyer   Buyer   `json:"buyer" gorm:"foreignKey:BuyerID"`
	Product Product `json:"product" gorm:"foreignKey:ProductID"`
}
