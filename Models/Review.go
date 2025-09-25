package models

import (
	"time"

	"gorm.io/gorm"
)

// Review represents product reviews - MVP
type Review struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	ProductID uint           `json:"product_id" gorm:"not null"`
	BuyerID   uint           `json:"buyer_id" gorm:"not null"`
	OrderID   *uint          `json:"order_id"` // Must be linked to actual purchase
	Rating    int            `json:"rating" gorm:"not null;check:rating >= 1 AND rating <= 5"`
	Title     string         `json:"title" gorm:"size:255"` // Added review title
	Comment   string         `json:"comment" gorm:"type:text"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	Product Product       `json:"product" gorm:"foreignKey:ProductID"`
	Buyer   Buyer         `json:"buyer" gorm:"foreignKey:BuyerID"`
	Order   Order         `json:"order" gorm:"foreignKey:OrderID"`
	Images  []ReviewImage `json:"images" gorm:"foreignKey:ReviewID"`
}
