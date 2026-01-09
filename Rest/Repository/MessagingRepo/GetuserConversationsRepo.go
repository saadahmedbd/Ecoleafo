package messagingrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *messagingRepositoryImpl) GetUserConversations(userID uint, userType string) ([]models.Conversation, error) {
	var convs []models.Conversation
	err := r.db.Where(
		"(participant1_id = ? AND participant1_type = ?) OR (participant2_id = ? AND participant2_type = ?)",
		userID, userType, userID, userType,
	).Preload("Product").Preload("Order").Order("last_message_at DESC NULLS LAST, created_at DESC").Find(&convs).Error
	return convs, err
}
