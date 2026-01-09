package messagingservice

import (
	"time"

	messagingdto "github.com/saadahmedbd/Treestore/Rest/DTO/MessagingDTO"
	messagingrepo "github.com/saadahmedbd/Treestore/Rest/Repository/MessagingRepo"
)

type MessagingService interface {
	CreateConversation(userCtx *messagingdto.UserContext, req *messagingdto.CreateConversationRequest) (*messagingdto.ConversationResponse, error)
	GetConversations(userCtx *messagingdto.UserContext) (*messagingdto.ConversationListResponse, error)
	GetConversationDetails(userCtx *messagingdto.UserContext, convID uint) (*messagingdto.ConversationResponse, error)
	SendMessage(userCtx *messagingdto.UserContext, req *messagingdto.SendMessageRequest) (*messagingdto.MessageResponse, error)
	GetMessages(userCtx *messagingdto.UserContext, convID uint, since time.Time, limit int) (*messagingdto.MessagesResponse, error)
	MarkAsRead(userCtx *messagingdto.UserContext, convID uint) error
	GetUnreadCount(userCtx *messagingdto.UserContext) (*messagingdto.UnreadCountResponse, error)
}

type messagingServiceImpl struct {
	messagerepo messagingrepo.MessagingRepository
}

func NewMessagingService(messagerepo messagingrepo.MessagingRepository) MessagingService {
	return &messagingServiceImpl{
		messagerepo: messagerepo,
	}
}
