package models

import (
	"time"

	"gorm.io/gorm"
)

type Conversation struct {
	ID               uint   `json:"id" gorm:"primaryKey"`
	Participant1ID   uint   `json:"participant1_id" gorm:"not null;index"`
	Participant1Type string `json:"participant1_type" gorm:"size:20;not null"`
	Participant2ID   uint   `json:"participant2_id" gorm:"not null;index"`
	Participant2Type string `json:"participant2_type" gorm:"size:20;not null"`

	ContextType string `json:"context_type" gorm:"size:20"`
	ContextID   *uint  `json:"context_id"`

	LastMessage   string     `json:"last_message" gorm:"type:text"`
	LastMessageAt *time.Time `json:"last_message_at"`

	Participant1UnreadCount int `json:"participant1_unread_count" gorm:"default:0"`
	Participant2UnreadCount int `json:"participant2_unread_count" gorm:"default:0"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Messages []Message `json:"messages,omitempty" gorm:"foreignKey:ConversationID"`
	Product  *Product  `json:"product,omitempty" gorm:"foreignKey:ContextID"`
	Order    *Order    `json:"order,omitempty" gorm:"foreignKey:ContextID"`
}

type Message struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	ConversationID uint   `json:"conversation_id" gorm:"not null;index"`
	SenderID       uint   `json:"sender_id" gorm:"not null;index"`
	SenderType     string `json:"sender_type" gorm:"size:20;not null"`

	MessageText string `json:"message_text" gorm:"type:text;not null"`

	IsRead bool       `json:"is_read" gorm:"default:false;index"`
	ReadAt *time.Time `json:"read_at"`

	CreatedAt time.Time      `json:"created_at" gorm:"index"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Conversation Conversation        `json:"conversation,omitempty" gorm:"foreignKey:ConversationID"`
	Attachments  []MessageAttachment `json:"attachments,omitempty" gorm:"foreignKey:MessageID"`
}

type MessageAttachment struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	MessageID uint   `json:"message_id" gorm:"not null;index"`
	FileURL   string `json:"file_url" gorm:"size:500;not null"`
	FileName  string `json:"file_name" gorm:"size:255"`
	FileType  string `json:"file_type" gorm:"size:50"`
	FileSize  int64  `json:"file_size"`

	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Message Message `json:"message,omitempty" gorm:"foreignKey:MessageID"`
}
