package models

import (
	"time"

	"gorm.io/gorm"
)

// InventoryThreshold allows flexible low-stock alert per product or category
type InventoryThreshold struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	SellerID    uint           `json:"seller_id" gorm:"not null;index"`
	ProductID   uint           `json:"product_id" gorm:"index"`
	CategoryID  uint           `json:"category_id" gorm:"index"`
	MinQuantity int            `json:"min_quantity" gorm:"not null;default:1"`
	AlertLevel  string         `json:"alert_level" gorm:"size:20;default:'normal'"` // e.g. low, critical
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relations
	Seller   User      `json:"seller" gorm:"foreignKey:SellerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product  *Product  `json:"product,omitempty" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Category *Category `json:"category,omitempty" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (InventoryThreshold) TableName() string {
	return "inventory_thresholds"
}
