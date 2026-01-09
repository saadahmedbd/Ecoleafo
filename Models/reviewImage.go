package models

import "time"

type ReviewImage struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	ReviewID  uint      `json:"review_id" gorm:"not null"`
	ImageURL  string    `json:"image_url" gorm:"size:500;not null"`
	AltText   string    `json:"alt_text" gorm:"size:255"`
	CreatedAt time.Time `json:"created_at"`

	Review Review `json:"review" gorm:"foreignKey:ReviewID"`
}
