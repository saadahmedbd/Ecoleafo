package messagingservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	messagingdto "github.com/saadahmedbd/Treestore/Rest/DTO/MessagingDTO"
)

func (s *messagingServiceImpl) buildContextInfo(contextType string, contextID uint, conv *models.Conversation) *messagingdto.ConversationContextInfo {
	switch contextType {
	case "product":
		if conv.Product != nil {
			contextInfo := &messagingdto.ConversationContextInfo{
				Type: "product",
				ID:   conv.Product.ID,
				Name: conv.Product.Name,
			}
			if len(conv.Product.Images) > 0 {
				contextInfo.Image = conv.Product.Images[0].ImageURL
			}
			return contextInfo
		}
	case "order":
		if conv.Order != nil {
			return &messagingdto.ConversationContextInfo{
				Type: "order",
				ID:   conv.Order.ID,
				Name: conv.Order.OrderNumber,
			}
		}
	case "support":
		return &messagingdto.ConversationContextInfo{
			Type: "support",
			ID:   contextID,
			Name: "Customer Support",
		}
	}
	return nil
}
