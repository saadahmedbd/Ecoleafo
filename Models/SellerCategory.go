package models

import "time"

type SellerCategory struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	SellerID   uint      `json:"seller_id" gorm:"not null;uniqueIndex:idx_seller_category"`
	CategoryID uint      `json:"category_id" gorm:"not null;uniqueIndex:idx_seller_category"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	Seller   User     `json:"seller" gorm:"foreignKey:SellerID"`
	Category Category `json:"category" gorm:"foreignKey:CategoryID"`
}
