package messagingdto

// Request DTOs
type CreateConversationRequest struct {
	RecipientID    uint   `json:"recipient_id" validate:"required"`
	RecipientType  string `json:"recipient_type" validate:"required,oneof=buyer seller admin"`
	ContextType    string `json:"context_type" validate:"omitempty,oneof=product order support"`
	ContextID      *uint  `json:"context_id"`
	InitialMessage string `json:"initial_message" validate:"required,min=1,max=1000"`
}
