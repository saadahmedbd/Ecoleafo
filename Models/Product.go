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
	CategoryID uint `json:"category_id" gorm:"not null"`

	//pricing
	Price           float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	DiscountPrice   float64 `json:"discount price" gorm:"type:decimal(10,2)"`
	DiscountPercent float64 `json:"discount_percent" gorm:"type:decimal(5,2)"`

	//tree specific attributes
	Height         string `json:"height" gorm:"not null;default:''"`
	Age            string `json:"age" gorm:"not null;default:''"`
	TreeType       string `json:"tree_type" gorm:"size:100"`       // fruit, ornamental, shade, medicinal
	PotSize        string `json:"pot_size" gorm:"size:50"`         // Added pot size
	ScientificName string `json:"scientific_name" gorm:"size:200"` // Added scientific name
	CommonNames    string `json:"common_names" gorm:"size:500"`

	//inventory logistics
	Quantity    int     `json:"quantity" gorm:"default:0"`
	MinQuantity int     `json:"min_quantity" gorm:"default:1"`   // Minimum order quantity
	Weight      float64 `json:"weight" gorm:"type:decimal(8,2)"` // For shipping calculation

	//product status
	IsActive   bool `json:"is_active" gorm:"default:true"`
	IsApproved bool `json:"is_approved" gorm:"default:false"` // Admin approval required
	IsFeatured bool `json:"is_featured" gorm:"default:false"` // Added featured products

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// SEO & Analytics (NEW)
	MetaTitle       string  `json:"meta_title" gorm:"size:255"`                        // SEO title
	MetaDescription string  `json:"meta_description" gorm:"size:500"`                  // SEO description
	ViewCount       int     `json:"view_count" gorm:"default:0"`                       // Product views
	SaleCount       int     `json:"sale_count" gorm:"default:0"`                       // Times sold
	AverageRating   float64 `json:"average_rating" gorm:"type:decimal(3,2);default:0"` // Average rating
	ReviewCount     int     `json:"review_count" gorm:"default:0"`                     // Number of reviews

	// Relationships
	Seller        User               `json:"seller" gorm:"foreignKey:SellerID"`
	Category      Category           `json:"category" gorm:"foreignKey:CategoryID"`
	Images        []ProductImage     `json:"images" gorm:"foreignKey:ProductID"`     // Multiple images
	Attributes    []ProductAttribute `json:"attributes" gorm:"foreignKey:ProductID"` // Flexible attributes
	CartItems     []CartItem         `json:"cart_items" gorm:"foreignKey:ProductID"`
	OrderItems    []OrderItem        `json:"order_items" gorm:"foreignKey:ProductID"`
	Reviews       []Review           `json:"reviews" gorm:"foreignKey:ProductID"`
	WishlistItems []Wishlist         `json:"wishlist_items" gorm:"foreignKey:ProductID"`
}
