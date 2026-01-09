package messagingservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	messagingdto "github.com/saadahmedbd/Treestore/Rest/DTO/MessagingDTO"
)

func (s *messagingServiceImpl) buildConversationResponse(conv *models.Conversation, userCtx *messagingdto.UserContext) (*messagingdto.ConversationResponse, error) {
	// Determine other user
	var otherUserID uint
	var otherUserType string
	var unreadCount int

	if conv.Participant1ID == userCtx.RoleID && conv.Participant1Type == userCtx.UserType {
		otherUserID = conv.Participant2ID
		otherUserType = conv.Participant2Type
		unreadCount = conv.Participant1UnreadCount
	} else {
		otherUserID = conv.Participant1ID
		otherUserType = conv.Participant1Type
		unreadCount = conv.Participant2UnreadCount
	}

	// Build other user info
	otherUserInfo := s.buildUserInfo(otherUserID, otherUserType)

	// Build context info
	var contextInfo *messagingdto.ConversationContextInfo
	if conv.ContextType != "" && conv.ContextID != nil {
		contextInfo = s.buildContextInfo(conv.ContextType, *conv.ContextID, conv)
	}

	return &messagingdto.ConversationResponse{
		ID:            conv.ID,
		OtherUser:     otherUserInfo,
		LastMessage:   conv.LastMessage,
		LastMessageAt: conv.LastMessageAt,
		UnreadCount:   unreadCount,
		Context:       contextInfo,
		CreatedAt:     conv.CreatedAt,
	}, nil
}
