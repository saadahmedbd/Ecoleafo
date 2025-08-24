package models

import (
	"fmt"

	"time"

	"gorm.io/gorm"
)

// User represents sellers who can manage products
type User struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	RoleID    uint   `json:"role_id" gorm:"not null"`
	FirstName string `json:"first_name" gorm:"size:50;not null"`
	LastName  string `json:"last_name" gorm:"size:50;not null"`
	Email     string `json:"email" gorm:"uniqueIndex;size:100;not null"`
	Password  string `json:"-" gorm:"size:255;not null"` // Hidden in JSON
	Phone     string `json:"phone" gorm:"size:20"`
	StoreName string `json:"store_name" gorm:"size:100"`         // Seller's store name
	StoreDesc string `json:"store_description" gorm:"type:text"` // Store description
	// Commission  float64        `json:"commission" gorm:"type:decimal(5,2);default:10"` // Commission percentage
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	IsVerified  bool           `json:"is_verified" gorm:"default:false"` // Seller verification
	TotalSales  float64        `json:"total_sales" gorm:"type:decimal(10,2);default:0"`
	TotalOrders int            `json:"total_orders" gorm:"default:0"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	Role     Role      `json:"role" gorm:"foreignKey:RoleID"`
	Products []Product `json:"products" gorm:"foreignKey:SellerID"` // Seller's products
}

func InsertData(db *gorm.DB, user1 User) error {
	result := db.Create(&user1)

	if result.Error != nil {
		return result.Error
	}
	fmt.Println("data inserted", user1)
	return nil
}
