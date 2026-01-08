package messagingservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	messagingdto "github.com/saadahmedbd/Treestore/Rest/DTO/MessagingDTO"
)

func (s *messagingServiceImpl) isParticipant(conv *models.Conversation, userCtx *messagingdto.UserContext) bool {
	return (conv.Participant1ID == userCtx.RoleID && conv.Participant1Type == userCtx.UserType) ||
		(conv.Participant2ID == userCtx.RoleID && conv.Participant2Type == userCtx.UserType)
}

func (s *messagingServiceImpl) getRecipientNumber(conv *models.Conversation, userCtx *messagingdto.UserContext) int {
	if conv.Participant1ID == userCtx.RoleID && conv.Participant1Type == userCtx.UserType {
		return 2
	}
	return 1
}
