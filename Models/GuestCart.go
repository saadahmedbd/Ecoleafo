package models

import "time"

type GuestCartItem struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	SessionID  string    `json:"session_id" gorm:"size:255;not null;index"`
	ProductID  uint      `json:"product_id" gorm:"not null"`
	Quantity   int       `json:"quantity" gorm:"not null"`
	Price      float64   `json:"price" gorm:"type:decimal(10,2);not null"`
	CreatyedAt time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	Product Product `json:"product" gorm:"foreignKey:ProductID"`
}
