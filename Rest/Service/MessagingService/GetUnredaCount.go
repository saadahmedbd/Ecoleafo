package messagingservice

import (
	"fmt"

	messagingdto "github.com/saadahmedbd/Treestore/Rest/DTO/MessagingDTO"
)

func (s *messagingServiceImpl) GetUnreadCount(userCtx *messagingdto.UserContext) (*messagingdto.UnreadCountResponse, error) {
	count, err := s.messagerepo.GetUnreadMessagesCount(userCtx.RoleID, userCtx.UserType)
	if err != nil {
		return nil, fmt.Errorf("failed to get unread count: %w", err)
	}

	return &messagingdto.UnreadCountResponse{
		TotalUnread: int(count),
	}, nil
}
