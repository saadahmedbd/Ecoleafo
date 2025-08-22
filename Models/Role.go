package models

import (
	"time"

	"gorm.io/gorm"
)

// Role represents user roles (admin, seller, buyer)
type Role struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"size:50;uniqueIndex;not null"` // admin, seller, buyer
	Description string         `json:"description" gorm:"size:255"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	Users  []User  `json:"users" gorm:"foreignKey:RoleID"`  // Sellers
	Buyers []Buyer `json:"buyers" gorm:"foreignKey:RoleID"` // Buyers
}
