package messagingservice

import (
	"errors"
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	messagingdto "github.com/saadahmedbd/Treestore/Rest/DTO/MessagingDTO"
	"gorm.io/gorm"
)

func (s *messagingServiceImpl) CreateConversation(userCtx *messagingdto.UserContext, req *messagingdto.CreateConversationRequest) (*messagingdto.ConversationResponse, error) {
	// Validate recipient exists
	if err := s.validateRecipient(req.RecipientID, req.RecipientType); err != nil {
		return nil, err
	}

	// Check if conversation already exists
	existingConv, err := s.messagerepo.GetConversationByParticipants(
		userCtx.RoleID, userCtx.UserType,
		req.RecipientID, req.RecipientType,
		req.ContextType, req.ContextID,
	)

	var conv *models.Conversation
	if err == nil && existingConv != nil {
		// Use existing conversation
		conv = existingConv
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		// Create new conversation
		conv = &models.Conversation{
			Participant1ID:          userCtx.RoleID,
			Participant1Type:        userCtx.UserType,
			Participant2ID:          req.RecipientID,
			Participant2Type:        req.RecipientType,
			ContextType:             req.ContextType,
			ContextID:               req.ContextID,
			Participant1UnreadCount: 0,
			Participant2UnreadCount: 0,
		}

		if err := s.messagerepo.CreateConversation(conv); err != nil {
			return nil, fmt.Errorf("failed to create conversation: %w", err)
		}
	} else {
		return nil, err
	}

	// Send initial message
	msgReq := &messagingdto.SendMessageRequest{
		ConversationID: conv.ID,
		MessageText:    req.InitialMessage,
	}

	if _, err := s.SendMessage(userCtx, msgReq); err != nil {
		return nil, fmt.Errorf("failed to send initial message: %w", err)
	}

	// Return conversation response
	return s.buildConversationResponse(conv, userCtx)
}
