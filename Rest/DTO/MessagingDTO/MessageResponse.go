package messagingdto

import "time"

type MessageResponse struct {
	ID          uint       `json:"id"`
	SenderID    uint       `json:"sender_id"`
	SenderName  string     `json:"sender_name"`
	SenderType  string     `json:"sender_type"`
	MessageText string     `json:"message_text"`
	IsRead      bool       `json:"is_read"`
	ReadAt      *time.Time `json:"read_at"`
	CreatedAt   time.Time  `json:"created_at"`
	IsMine      bool       `json:"is_mine"`
}

type MessagesResponse struct {
	Messages     []MessageResponse     `json:"messages"`
	HasMore      bool                  `json:"has_more"`
	Conversation *ConversationResponse `json:"conversation,omitempty"`
}

type UnreadCountResponse struct {
	TotalUnread int `json:"total_unread"`
}

type UserContext struct {
	UserID   uint   `json:"user_id"`
	Email    string `json:"email"`
	UserType string `json:"user_type"`
	RoleID   uint   `json:"role_id"`
}
