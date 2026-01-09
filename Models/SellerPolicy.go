package models

import (
	"time"

	"gorm.io/gorm"
)

type SellerPolicy struct {
	ID             uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	SellerID       uint   `json:"seller_id" gorm:"uniqueIndex;not null"`
	ReturnPolicy   string `json:"return_policy" gorm:"type:text"`
	ShippingPolicy string `json:"shipping_policy" gorm:"type:text"`
	FAQ            string `json:"faq" gorm:"type:text"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	Seller User `json:"seller" gorm:"foreignKey:SellerID;references:ID"`
}

func (SellerPolicy) TableName() string {
	return "seller_policies"
}
