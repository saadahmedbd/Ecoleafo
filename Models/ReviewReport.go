package models

import "time"

// Review Report (when users flag reviews)
type ReviewReport struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	ReviewID   uint      `json:"review_id" gorm:"not null;index"`
	ReporterID uint      `json:"reporter_id"`                     // Can be buyer or seller
	Reason     string    `json:"reason" gorm:"size:100;not null"` // spam, inappropriate, fake, etc
	Details    string    `json:"details" gorm:"type:text"`
	Status     string    `json:"status" gorm:"size:20;default:'pending'"` // pending, reviewed, dismissed
	CreatedAt  time.Time `json:"created_at"`

	// Relationships
	Review Review `json:"review" gorm:"foreignKey:ReviewID;references:ID"`
}
