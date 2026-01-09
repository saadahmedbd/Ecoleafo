package messagingrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *messagingRepositoryImpl) GetConversationByID(id uint) (*models.Conversation, error) {
	var conv models.Conversation
	err := r.db.Preload("Product").Preload("Order").First(&conv, id).Error
	return &conv, err
}
