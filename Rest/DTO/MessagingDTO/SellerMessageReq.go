// The `package messagingdto` statement in Go is declaring that the following code belongs to the
// `messagingdto` package. This package is used to organize related code and provide a namespace for
// the types and functions defined within it. It helps in structuring and managing code in a modular
// and reusable way.
package messagingdto

type SendMessageRequest struct {
	ConversationID uint   `json:"conversation_id" validate:"required"`
	MessageText    string `json:"message_text" validate:"required,min=1,max=5000"`
}
