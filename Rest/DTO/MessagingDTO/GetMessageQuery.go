package messagingdto

import "time"

type GetMessagesQuery struct {
	ConversationID uint      `json:"conversation_id" validate:"required"`
	Since          time.Time `json:"since"`
	Limit          int       `json:"limit" validate:"omitempty,min=1,max=100"`
}

type MarkAsReadRequest struct {
	ConversationID uint `json:"conversation_id" validate:"required"`
}
