package models

import (
	"time"

	"gorm.io/gorm"
)

type AdminInvitation struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	Email      string         `json:"email" gorm:"size:100;uniqueIndex;not null"`
	Token      string         `json:"token" gorm:"size:255;uniqueIndex;not null"`
	Role       string         `json:"role" gorm:"size:20;not null;default:'admin'"` // admin, super_admin
	Department string         `json:"department" gorm:"size:50"`
	InvitedBy  uint           `json:"invited_by" gorm:"not null"` // FK -> Admin.ID
	IsUsed     bool           `json:"is_used" gorm:"default:false"`
	ExpiresAt  time.Time      `json:"expires_at" gorm:"not null"`
	CreatedAt  time.Time      `json:"created_at"`
	UsedAt     *time.Time     `json:"used_at,omitempty"` // track when invite was used
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	InvitedByAdmin Admin `json:"invited_by_admin" gorm:"foreignKey:InvitedBy;references:ID"`
}
