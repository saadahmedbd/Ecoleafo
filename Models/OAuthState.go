package models

import "time"

type OAuthState struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	State     string    `gorm:"uniqueIndex;not null" json:"state"`
	Role      string    `gorm:"not null" json:"role"`
	ExpiresAt time.Time `gorm:"not null;index" json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}
