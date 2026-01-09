package models

import "time"

type SellerPaymentMethod struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	SellerID      uint      `json:"seller_id" gorm:"not null"`
	Type          string    `json:"type" gorm:"size:50;not null"` // bank_transfer, bkash, nagad, rocket
	AccountName   string    `json:"account_name" gorm:"size:100;not null"`
	AccountNumber string    `json:"account_number" gorm:"size:100;not null"`
	BankName      string    `json:"bank_name" gorm:"size:100"`
	BankCode      string    `json:"bank_code" gorm:"size:20"`
	RoutingNumber string    `json:"routing_number" gorm:"size:50"`
	IsDefault     bool      `json:"is_default" gorm:"default:false"`
	IsActive      bool      `json:"is_active" gorm:"default:true"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	Seller User `json:"seller" gorm:"foreignKey:SellerID"`
}
