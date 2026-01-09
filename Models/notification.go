package models

import "time"

type Notification struct {
	ID        uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint       `json:"user_id" gorm:"not null"`
	Type      string     `json:"type" gorm:"size:50;not null"`
	Title     string     `json:"title" gorm:"size:255;not null"`
	Message   string     `json:"message" gorm:"type:text;not null"`
	Data      string     `json:"data" gorm:"type:json"` // Additional data
	IsRead    bool       `json:"is_read" gorm:"default:false"`
	ReadAt    *time.Time `json:"read_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
