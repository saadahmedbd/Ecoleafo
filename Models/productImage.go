package models

import "time"

type ProductImage struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductID uint      `json:"product_id" gorm:"not null"`
	ImageURL  string    `json:"image_url" gorm:"size:500;not null"`
	AltText   string    `json:"alt_text" gorm:"size:255"`
	IsPrimary bool      `json:"is_primary" gorm:"default:false"`
	SortOrder int       `json:"sort_order" gorm:"default:0"`
	Type      string    `json:"type" gorm:"size:50;default:'gallery'"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Product Product `json:"-" gorm:"foreignKey:ProductID"`
}
