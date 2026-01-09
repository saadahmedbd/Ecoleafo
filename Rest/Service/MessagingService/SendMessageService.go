package messagingservice

import (
	"errors"
	"fmt"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	messagingdto "github.com/saadahmedbd/Treestore/Rest/DTO/MessagingDTO"
)

func (s *messagingServiceImpl) SendMessage(userCtx *messagingdto.UserContext, req *messagingdto.SendMessageRequest) (*messagingdto.MessageResponse, error) {
	// Get conversation
	conv, err := s.messagerepo.GetConversationByID(req.ConversationID)
	if err != nil {
		return nil, fmt.Errorf("conversation not found: %w", err)
	}

	// Verify user is participant
	if !s.isParticipant(conv, userCtx) {
		return nil, errors.New("unauthorized: not a participant")
	}

	// Create message
	msg := &models.Message{
		ConversationID: req.ConversationID,
		SenderID:       userCtx.RoleID,
		SenderType:     userCtx.UserType,
		MessageText:    req.MessageText,
		IsRead:         false,
	}

	if err := s.messagerepo.CreateMessage(msg); err != nil {
		return nil, fmt.Errorf("failed to create message: %w", err)
	}

	// Update conversation
	now := time.Now()
	conv.LastMessage = req.MessageText
	conv.LastMessageAt = &now

	// Increment unread count for recipient
	recipientNum := s.getRecipientNumber(conv, userCtx)
	s.messagerepo.IncrementUnreadCount(conv.ID, recipientNum)

	if err := s.messagerepo.UpdateConversation(conv); err != nil {
		return nil, fmt.Errorf("failed to update conversation: %w", err)
	}

	// Build response
	senderName := s.getUserName(userCtx.RoleID, userCtx.UserType)

	return &messagingdto.MessageResponse{
		ID:          msg.ID,
		SenderID:    msg.SenderID,
		SenderName:  senderName,
		SenderType:  msg.SenderType,
		MessageText: msg.MessageText,
		IsRead:      msg.IsRead,
		ReadAt:      msg.ReadAt,
		CreatedAt:   msg.CreatedAt,
		IsMine:      true,
	}, nil
}
