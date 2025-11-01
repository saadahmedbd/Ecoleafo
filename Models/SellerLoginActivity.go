package models

import (
	"time"

	"gorm.io/gorm"
)

type SellerLoginActivity struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	SellerID  uint   `json:"seller_id" gorm:"not null;index"` // References User.ID
	Device    string `json:"device" gorm:"size:100"`          // e.g., "Chrome on MacOS"
	Browser   string `json:"browser" gorm:"size:50"`          // e.g., "Chrome 120.0"
	OS        string `json:"os" gorm:"size:50"`               // e.g., "MacOS"
	Location  string `json:"location" gorm:"size:100"`        // e.g., "Dhaka, Bangladesh"
	IPAddress string `json:"ip_address" gorm:"size:45"`       // IPv4 or IPv6
	IsCurrent bool   `json:"is_current" gorm:"default:false"` // Current session

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	Seller User `json:"seller" gorm:"foreignKey:SellerID;references:ID"`
}

func (SellerLoginActivity) TableName() string {
	return "seller_login_activities"
}
