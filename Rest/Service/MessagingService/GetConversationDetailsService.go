package messagingservice

import (
	"errors"
	"fmt"

	messagingdto "github.com/saadahmedbd/Treestore/Rest/DTO/MessagingDTO"
)

func (s *messagingServiceImpl) GetConversationDetails(userCtx *messagingdto.UserContext, convID uint) (*messagingdto.ConversationResponse, error) {
	conv, err := s.messagerepo.GetConversationByID(convID)
	if err != nil {
		return nil, fmt.Errorf("conversation not found: %w", err)
	}

	// Verify user is participant
	if !s.isParticipant(conv, userCtx) {
		return nil, errors.New("unauthorized: not a participant")
	}

	return s.buildConversationResponse(conv, userCtx)
}
