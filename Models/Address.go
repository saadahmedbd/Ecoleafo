package models

import "time"

type Address struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	BuyerID      uint      `json:"buyer_id" gorm:"not null"`
	Type         string    `json:"type" gorm:"size:20;not null;default:'shipping'"` // shipping, billing
	FirstName    string    `json:"first_name" gorm:"size:50;not null"`
	LastName     string    `json:"last_name" gorm:"size:50;not null"`
	Company      string    `json:"company" gorm:"size:100"`
	AddressLine1 string    `json:"address_line_1" gorm:"size:255;not null"`
	AddressLine2 string    `json:"address_line_2" gorm:"size:255"`
	City         string    `json:"city" gorm:"size:50;not null"`
	State        string    `json:"state" gorm:"size:50;not null"`
	District     string    `json:"district" gorm:"size:50;not null"`
	PostalCode   string    `json:"postal_code" gorm:"size:20;not null"`
	Country      string    `json:"country" gorm:"size:50;not null;default:'Bangladesh'"`
	Phone        string    `json:"phone" gorm:"size:20"`
	IsDefault    bool      `json:"is_default" gorm:"default:false"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Buyer Buyer `json:"buyer" gorm:"foreignKey:BuyerID"`
}
