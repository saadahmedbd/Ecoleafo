package messagingdto

type SendMessageRequest struct {
	ConversationID uint   `json:"conversation_id" validate:"required"`
	MessageText    string `json:"message_text" validate:"required,min=1,max=5000"`
}
