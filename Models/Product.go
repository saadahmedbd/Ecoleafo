package models

import (
	"time"

	"gorm.io/gorm"
)

// Product represents products sold by sellers - MVP version

type Product struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	SellerID    uint   `json:"seller_id" gorm:"not null"` // Which seller owns this product
	Name        string `json:"name" gorm:"size:255;not null"`
	Slug        string `json:"slug" gorm:"size:255;uniqueIndex;not null"` //a slug is the user-friendly, readable, and SEO-optimized part of a URL that identifies a specific product, category, page, or post.
	Description string `json:"description" gorm:"type:text"`
	SKU         string `json:"sku" gorm:"size:100;uniqueIndex;not null"` // SKU stands for Stock Keeping Unit.
	// It’s a unique identifier assigned to each product or product variant to track inventory, pricing, and sales.
	CategoryID    uint           `json:"category_id" gorm:"not null"`
	Price         float64        `json:"price" gorm:"type:decimal(10,2);not null"`
	DiscountPrice float64        `json:"discount price" gorm:"type:decimal(10,2)"`
	Height        string         `json:"height" gorm:"not null;default:''"`
	Age           string         `json:"age" gorm:"not null;default:''"`
	Quantity      int            `json:"quantity" gorm:"default:0"`
	Image         string         `json:"image" gorm:"size:500"` // Single image only
	IsActive      bool           `json:"is_active" gorm:"default:true"`
	IsApproved    bool           `json:"is_approved" gorm:"default:false"` // Admin approval required
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	Seller     User        `json:"seller" gorm:"foreignKey:SellerID"`
	Category   Category    `json:"category" gorm:"foreignKey:CategoryID"`
	CartItems  []CartItem  `json:"cart_items" gorm:"foreignKey:ProductID"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:ProductID"`
	Reviews    []Review    `json:"reviews" gorm:"foreignKey:ProductID"`
}
