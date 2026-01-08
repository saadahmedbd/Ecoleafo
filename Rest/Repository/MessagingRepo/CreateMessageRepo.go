package messagingrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *messagingRepositoryImpl) CreateMessage(msg *models.Message) error {
	return r.db.Create(msg).Error
}
