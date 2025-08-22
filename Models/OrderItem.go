package models

import "time"

// OrderItem represents individual items in an order - MVP
type OrderItem struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	OrderID       uint      `json:"order_id" gorm:"not null"`
	ProductID     uint      `json:"product_id" gorm:"not null"`
	SellerID      uint      `json:"seller_id" gorm:"not null"` // Track which seller gets the sale
	Quantity      int       `json:"quantity" gorm:"not null;check:quantity > 0"`
	Price         float64   `json:"price" gorm:"type:decimal(10,2);not null"`
	Total         float64   `json:"total" gorm:"type:decimal(10,2);not null"`
	Commission    float64   `json:"commission" gorm:"type:decimal(10,2);not null"`     // Platform commission
	SellerEarning float64   `json:"seller_earning" gorm:"type:decimal(10,2);not null"` // Seller's earning after commission
	CreatedAt     time.Time `json:"created_at"`

	// Relationships
	Order   Order   `json:"order" gorm:"foreignKey:OrderID"`
	Product Product `json:"product" gorm:"foreignKey:ProductID"`
	Seller  User    `json:"seller" gorm:"foreignKey:SellerID"`
}
