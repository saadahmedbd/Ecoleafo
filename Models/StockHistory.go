package models

import (
	"time"

	"gorm.io/gorm"
)

type StockHistory struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductID   uint      `json:"product_id" gorm:"not null;index"`
	SellerID    uint      `json:"seller_id" gorm:"not null;index"`
	Type        string    `json:"type" gorm:"size:20;not null"` // restock, sale, adjustment, return
	Quantity    int       `json:"quantity" gorm:"not null"`
	PrevStock   int       `json:"prev_stock" gorm:"not null"`
	NewStock    int       `json:"new_stock" gorm:"not null"`
	Reason      string    `json:"reason" gorm:"size:255"`
	Reference   string    `json:"reference" gorm:"size:100"` // order_id, adjustment_id, etc.
	CreatedBy   uint      `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	Product Product `json:"product" gorm:"foreignKey:ProductID"`
	Seller  User    `json:"seller" gorm:"foreignKey:SellerID"`
}

func (StockHistory) TableName() string {
	return "stock_histories"
}
