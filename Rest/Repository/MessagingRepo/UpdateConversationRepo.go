package messagingrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *messagingRepositoryImpl) UpdateConversation(conv *models.Conversation) error {
	return r.db.Save(conv).Error
}
