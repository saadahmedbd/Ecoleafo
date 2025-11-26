package models

import "time"

type AuditLog struct {
	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`

	// Actor Information
	ActorID    *uint  `json:"actor_id" gorm:"index"`
	ActorType  string `json:"actor_type" gorm:"size:20;index"` // admin, seller, buyer, system
	ActorName  string `json:"actor_name" gorm:"size:100"`
	ActorEmail string `json:"actor_email" gorm:"size:100"`

	// Action Details
	Action      string `json:"action" gorm:"size:100;not null;index"` // create, update, delete, view, login, etc.
	ActionGroup string `json:"action_group" gorm:"size:50;index"`     // user, product, order, payment, auth, system
	Description string `json:"description" gorm:"type:text"`

	// Entity Information
	EntityType string `json:"entity_type" gorm:"size:50;index"` // user, seller, product, order, review, etc.
	EntityID   *uint  `json:"entity_id" gorm:"index"`
	EntityName string `json:"entity_name" gorm:"size:255"`

	// Change Tracking
	OldValues string `json:"old_values" gorm:"type:json"` // Previous state
	NewValues string `json:"new_values" gorm:"type:json"` // New state
	Changes   string `json:"changes" gorm:"type:json"`    // What changed (diff)

	// Request Information
	IPAddress     string `json:"ip_address" gorm:"size:45"`
	UserAgent     string `json:"user_agent" gorm:"size:500"`
	RequestURL    string `json:"request_url" gorm:"size:500"`
	RequestMethod string `json:"request_method" gorm:"size:10"`

	// Status & Classification
	Status   string `json:"status" gorm:"size:20;default:'success'"` // success, failed, warning
	Severity string `json:"severity" gorm:"size:20;default:'info'"`  // info, warning, critical
	Category string `json:"category" gorm:"size:50;index"`           // security, business, system

	// Metadata
	Metadata string `json:"metadata" gorm:"type:json"` // Additional context

	CreatedAt time.Time `json:"created_at" gorm:"index"`
}
