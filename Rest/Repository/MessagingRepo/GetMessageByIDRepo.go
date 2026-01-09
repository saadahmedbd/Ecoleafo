package messagingrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *messagingRepositoryImpl) GetMessageByID(id uint) (*models.Message, error) {
	var msg models.Message
	err := r.db.First(&msg, id).Error
	return &msg, err
}
