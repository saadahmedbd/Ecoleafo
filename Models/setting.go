package models

import "time"

type Setting struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Key         string    `json:"key" gorm:"size:100;uniqueIndex;not null"`
	Value       string    `json:"value" gorm:"type:text"`
	Description string    `json:"description" gorm:"size:500"`
	Type        string    `json:"type" gorm:"size:20;default:'string'"` // string, number, boolean, json
	IsPublic    bool      `json:"is_public" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
