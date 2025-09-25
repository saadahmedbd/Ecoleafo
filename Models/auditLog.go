package models

import "time"

type AuditLog struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    *uint     `json:"user_id"`
	UserType  string    `json:"user_type" gorm:"size:20"` // seller, buyer, admin
	Action    string    `json:"action" gorm:"size:100;not null"`
	TableName string    `json:"table_name" gorm:"size:50;not null"`
	RecordID  uint      `json:"record_id"`
	OldValues string    `json:"old_values" gorm:"type:json"`
	NewValues string    `json:"new_values" gorm:"type:json"`
	IPAddress string    `json:"ip_address" gorm:"size:45"`
	UserAgent string    `json:"user_agent" gorm:"size:500"`
	CreatedAt time.Time `json:"created_at"`
}
