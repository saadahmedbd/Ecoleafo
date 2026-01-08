package messagingrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *messagingRepositoryImpl) ResetUnreadCount(convID uint, participantNumber int) error {
	field := "participant1_unread_count"
	if participantNumber == 2 {
		field = "participant2_unread_count"
	}
	return r.db.Model(&models.Conversation{}).Where("id = ?", convID).
		Update(field, 0).Error
}
