package models

import (
	"time"

	"gorm.io/gorm"
)

// Buyer represents customers who can place orders
type Buyer struct {
	ID                     uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	RoleID                 uint   `json:"role_id" gorm:"not null"`
	UserId                 uint   `json:"user_id" gorm:"uniqueindex"`
	Password               string `json:"-" gorm:"size:255;not null"` // Hidden in JSON
	Phone                  string `json:"phone" gorm:"size:20"`
	ProfilePicture         string `json:"profile_picture" gorm:"size=500"`
	ProfilePictureUrl      string `json:"profile_picture_url" gorm:"size=500"`
	ProfilePicturePublicID string `json:"profile_picture_public_id" gorm:"size=255"`
	Status                 string `json:"status" gorm:"size:20;default:'active'"`
	//address
	DefaultAddress   string `json:"default_address" gorm:"type:text"`
	DefaultAddressID *uint  `json:"default_address_id"` // Link to Address table

	IsActive      bool `json:"is_active" gorm:"default:true"`
	EmailVerified bool `json:"email_verified" gorm:"default:false"`

	LastOrderAt      *time.Time `json:"last_order_at"`
	TotalOrdersCount int        `json:"total_orders_count" gorm:"default:0"`
	TotalSpent       float64    `json:"total_spent" gorm:"type:decimal(10,2);default:0"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	Role        Role       `json:"role" gorm:"foreignKey:RoleID"`
	RegUser     *RegUser   `json:"reg_user" gorm:"foreignKey:UserId;references:ID"`
	DefaultAddr *Address   `json:"default_addr,omitempty" gorm:"foreignKey:DefaultAddressID;constraint:OnDelete:SET NULL;-:migration"` // Skip FK in migration
	Orders      []Order    `json:"orders" gorm:"foreignKey:BuyerID"`
	CartItems   []CartItem `json:"cart_items" gorm:"foreignKey:BuyerID"`
	Reviews     []Review   `json:"reviews" gorm:"foreignKey:BuyerID"`
	Wishlists   []Wishlist `json:"wishlists" gorm:"foreignKey:BuyerID"` // Added wishlist
	Addresses   []Address  `json:"addresses" gorm:"foreignKey:BuyerID"` // Added multiple addresses
}
