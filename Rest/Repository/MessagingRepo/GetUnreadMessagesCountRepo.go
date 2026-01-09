package messagingrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *messagingRepositoryImpl) GetUnreadMessagesCount(userID uint, userType string) (int64, error) {
	var count int64

	// Get all conversations for this user
	var convIDs []uint
	err := r.db.Model(&models.Conversation{}).
		Select("id").
		Where("(participant1_id = ? AND participant1_type = ?) OR (participant2_id = ? AND participant2_type = ?)",
			userID, userType, userID, userType).
		Pluck("id", &convIDs).Error

	if err != nil {
		return 0, err
	}

	// Count unread messages in these conversations where user is NOT the sender
	err = r.db.Model(&models.Message{}).
		Where("conversation_id IN ? AND is_read = ? AND NOT (sender_id = ? AND sender_type = ?)",
			convIDs, false, userID, userType).
		Count(&count).Error

	return count, err
}
