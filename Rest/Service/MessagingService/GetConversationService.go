package messagingservice

import (
	"fmt"

	messagingdto "github.com/saadahmedbd/Treestore/Rest/DTO/MessagingDTO"
)

func (s *messagingServiceImpl) GetConversations(userCtx *messagingdto.UserContext) (*messagingdto.ConversationListResponse, error) {
	// Use RoleID for sellers (users.id), UserID for buyers/admins
	queryID := userCtx.UserID
	if userCtx.UserType == "seller" {
		queryID = userCtx.RoleID
	}
	
	fmt.Printf("[DEBUG] GetConversations - UserID: %d, RoleID: %d, UserType: %s, QueryID: %d\n", 
		userCtx.UserID, userCtx.RoleID, userCtx.UserType, queryID)
	
	convs, err := s.messagerepo.GetUserConversations(queryID, userCtx.UserType)
	if err != nil {
		return nil, fmt.Errorf("failed to get conversations: %w", err)
	}
	
	fmt.Printf("[DEBUG] Found %d conversations\n", len(convs))

	responses := make([]messagingdto.ConversationResponse, 0, len(convs))
	totalUnread := 0

	for _, conv := range convs {
		resp, err := s.buildConversationResponse(&conv, userCtx)
		if err != nil {
			continue // Skip if error building response
		}
		responses = append(responses, *resp)
		totalUnread += resp.UnreadCount
	}

	return &messagingdto.ConversationListResponse{
		Conversations: responses,
		TotalUnread:   totalUnread,
	}, nil
}
