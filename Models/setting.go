package models

import (
	"time"

	"gorm.io/gorm"
)

type Setting struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Key         string         `json:"key" gorm:"size:100;uniqueIndex;not null"`
	Value       string         `json:"value" gorm:"type:text"`
	Description string         `json:"description" gorm:"size:500"`
	Type        string         `json:"type" gorm:"size:20;default:'string'"` // string, number, boolean, json
	IsPublic    bool           `json:"is_public" gorm:"default:false"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	//relationship
	UpdatedByAdminID uint   `json:"updated_by_admin_id"`        // Foreign key column
	UpdatedByAdmin   *Admin `json:"updated_by_admin,omitempty"` // Relation reference

}
