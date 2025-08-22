package models

import (
	"time"

	"gorm.io/gorm"
)

// // Order represents buyer orders - MVP
type Order struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	OrderNumber     string         `json:"order_number" gorm:"size:50;uniqueIndex;not null"`
	BuyerID         uint           `json:"buyer_id" gorm:"not null"`
	Status          string         `json:"status" gorm:"size:20;not null;default:'pending'"` // pending, confirmed, shipped, delivered, cancelled
	Total           float64        `json:"total" gorm:"type:decimal(10,2);not null"`
	PaymentStatus   string         `json:"payment_status" gorm:"size:20;not null;default:'pending'"` // pending, paid, failed
	ShippingAddress string         `json:"shipping_address" gorm:"type:text;not null"`
	CustomerEmail   string         `json:"customer_email" gorm:"size:100;not null"`
	CustomerPhone   string         `json:"customer_phone" gorm:"size:20"`
	Notes           string         `json:"notes" gorm:"type:text"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	Buyer      Buyer       `json:"buyer" gorm:"foreignKey:BuyerID"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
}
