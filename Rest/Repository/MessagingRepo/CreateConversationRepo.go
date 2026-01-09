package messagingrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *messagingRepositoryImpl) CreateConversation(conv *models.Conversation) error {
	return r.db.Create(conv).Error
}
