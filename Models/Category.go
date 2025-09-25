package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"size:100;not null"`
	Slug        string         `json:"slug" gorm:"size:100;uniqueIndex;not null"`
	Description string         `json:"description" gorm:"type:text"`
	Image       string         `json:"image" gorm:"size:500"` // Category image
	Icon        string         `json:"icon" gorm:"size:100"`  // Category icon
	ParentID    *uint          `json:"parent_id"`             // For nested categories
	SortOrder   int            `json:"sort_order" gorm:"default:0"`
	IsFeatured  bool           `json:"is_featured" gorm:"default:false"` // Featured categories
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	Parent   *Category  `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Children []Category `json:"children" gorm:"foreignKey:ParentID"`
	Products []Product  `json:"products" gorm:"foreignKey:CategoryID"`
}
