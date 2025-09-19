package models

import (
	"time"

	"gorm.io/gorm"
)

// Buyer represents customers who can place orders
type Buyer struct {
	ID               uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	RoleID           uint           `json:"role_id" gorm:"not null"`
	UserId           uint           `json:"user_id" gorm:"uniqueindex"`
	FirstName        string         `json:"first_name" gorm:"size:50;not null"`
	LastName         string         `json:"last_name" gorm:"size:50;not null"`
	Email            string         `json:"email" gorm:"uniqueIndex;size:100;not null"`
	Password         string         `json:"-" gorm:"size:255;not null"` // Hidden in JSON
	Phone            string         `json:"phone" gorm:"size:20"`
	DefaultAddress   string         `json:"default_address" gorm:"type:text"`
	IsActive         bool           `json:"is_active" gorm:"default:true"`
	EmailVerified    bool           `json:"email_verified" gorm:"default:false"`
	LastOrderAt      *time.Time     `json:"last_order_at"`
	TotalOrdersCount int            `json:"total_orders_count" gorm:"default:0"`
	TotalSpent       float64        `json:"total_spent" gorm:"type:decimal(10,2);default:0"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	Role      Role       `json:"role" gorm:"foreignKey:RoleID"`
	Orders    []Order    `json:"orders" gorm:"foreignKey:BuyerID"`
	CartItems []CartItem `json:"cart_items" gorm:"foreignKey:BuyerID"`
	Reviews   []Review   `json:"reviews" gorm:"foreignKey:BuyerID"`
}
