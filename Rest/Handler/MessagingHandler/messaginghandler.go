package messaginghandler

import (
	"github.com/go-playground/validator/v10"
	messagingservice "github.com/saadahmedbd/Treestore/Rest/Service/MessagingService"
)

type MessagingHandler struct {
	messageservice messagingservice.MessagingService
	validate       *validator.Validate
}

func NewMessagingHandler(messageservice messagingservice.MessagingService) *MessagingHandler {
	return &MessagingHandler{
		messageservice: messageservice,
		validate:       validator.New(),
	}
}
