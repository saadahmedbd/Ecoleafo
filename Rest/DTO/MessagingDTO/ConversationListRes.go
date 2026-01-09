package messagingdto

import "time"

// Response DTOs
type ConversationListResponse struct {
	Conversations []ConversationResponse `json:"conversations"`
	TotalUnread   int                    `json:"total_unread"`
}

type ConversationResponse struct {
	ID            uint                     `json:"id"`
	OtherUser     ConversationUserInfo     `json:"other_user"`
	LastMessage   string                   `json:"last_message"`
	LastMessageAt *time.Time               `json:"last_message_at"`
	UnreadCount   int                      `json:"unread_count"`
	Context       *ConversationContextInfo `json:"context,omitempty"`
	CreatedAt     time.Time                `json:"created_at"`
}

type ConversationUserInfo struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	UserType       string `json:"user_type"`
	ProfilePicture string `json:"profile_picture,omitempty"`
	StoreName      string `json:"store_name,omitempty"`
}
type ConversationContextInfo struct {
	Type  string `json:"type"`
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image,omitempty"`
}
