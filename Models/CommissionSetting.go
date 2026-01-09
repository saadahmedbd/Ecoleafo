package models

import "time"

// ============================================================================
// COMMISSION SETTINGS - Platform-wide commission configuration
// ============================================================================
type CommissionSetting struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	DefaultRate float64   `json:"default_rate" gorm:"type:decimal(5,2);not null;default:10.00"` // Default 10%
	Description string    `json:"description" gorm:"type:text"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	UpdatedBy   uint      `json:"updated_by"` // Admin who updated
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relationships
	UpdatedByAdmin *Admin `json:"updated_by_admin,omitempty" gorm:"foreignKey:UpdatedBy;references:ID"`
}
