package messagingservice

import (
	"errors"
	"fmt"
	"time"

	messagingdto "github.com/saadahmedbd/Treestore/Rest/DTO/MessagingDTO"
)

func (s *messagingServiceImpl) GetMessages(userCtx *messagingdto.UserContext, convID uint, since time.Time, limit int) (*messagingdto.MessagesResponse, error) {
	// Get conversation
	conv, err := s.messagerepo.GetConversationByID(convID)
	if err != nil {
		return nil, fmt.Errorf("conversation not found: %w", err)
	}

	// Verify user is participant
	if !s.isParticipant(conv, userCtx) {
		return nil, errors.New("unauthorized: not a participant")
	}

	// Get messages
	msgs, err := s.messagerepo.GetMessagesByConversationID(convID, since, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get messages: %w", err)
	}

	// Build responses
	responses := make([]messagingdto.MessageResponse, 0, len(msgs))
	for _, msg := range msgs {
		senderName := s.getUserName(msg.SenderID, msg.SenderType)
		isMine := msg.SenderID == userCtx.UserID && msg.SenderType == userCtx.UserType

		responses = append(responses, messagingdto.MessageResponse{
			ID:          msg.ID,
			SenderID:    msg.SenderID,
			SenderName:  senderName,
			SenderType:  msg.SenderType,
			MessageText: msg.MessageText,
			IsRead:      msg.IsRead,
			ReadAt:      msg.ReadAt,
			CreatedAt:   msg.CreatedAt,
			IsMine:      isMine,
		})
	}

	convResp, _ := s.buildConversationResponse(conv, userCtx)

	return &messagingdto.MessagesResponse{
		Messages:     responses,
		HasMore:      len(msgs) == limit,
		Conversation: convResp,
	}, nil
}
