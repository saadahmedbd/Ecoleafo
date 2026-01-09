package models

import "time"

// CartItem represents items in buyer's shopping cart - MVP
type CartItem struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	BuyerID   uint    `json:"buyer_id" gorm:"not null"`
	ProductID uint    `json:"product_id" gorm:"not null"`
	Quantity  int     `json:"quantity" gorm:"not null;check:quantity > 0"`
	Price     float64 `json:"price" gorm:"type:decimal(10,2);not null"`

	// Additional useful fields
	IsSavedForLater bool   `json:"is_saved_for_later" gorm:"default:false"` // Save for later feature
	IsSelected      bool   `json:"is_selected" gorm:"default:true"`         // Selected for checkout
	IsGift          bool   `json:"is_gift" gorm:"default:false"`            // Gift wrapping option
	GiftMessage     string `json:"gift_message" gorm:"size:500"`
	AddedFrom       string `json:"added_from" gorm:"size:50"` // web, mobile, wishlist

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relationships
	Buyer   Buyer   `json:"buyer" gorm:"foreignKey:BuyerID"`
	Product Product `json:"product" gorm:"foreignKey:ProductID"`
}

// Add index for better performance
func (CartItem) TableName() string {
	return "cart_items"
}
