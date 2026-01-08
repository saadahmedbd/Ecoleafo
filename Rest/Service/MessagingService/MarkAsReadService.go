package messagingservice

import (
	"errors"
	"fmt"

	messagingdto "github.com/saadahmedbd/Treestore/Rest/DTO/MessagingDTO"
)

func (s *messagingServiceImpl) MarkAsRead(userCtx *messagingdto.UserContext, convID uint) error {
	// Get conversation
	conv, err := s.messagerepo.GetConversationByID(convID)
	if err != nil {
		return fmt.Errorf("conversation not found: %w", err)
	}

	// Verify user is participant
	if !s.isParticipant(conv, userCtx) {
		return errors.New("unauthorized: not a participant")
	}

	// Mark messages as read
	if err := s.messagerepo.MarkMessagesAsRead(convID, userCtx.RoleID, userCtx.UserType); err != nil {
		return fmt.Errorf("failed to mark messages as read: %w", err)
	}

	// Reset unread count
	participantNum := 1
	if conv.Participant2ID == userCtx.RoleID && conv.Participant2Type == userCtx.UserType {
		participantNum = 2
	}

	return s.messagerepo.ResetUnreadCount(convID, participantNum)
}
