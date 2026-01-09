package messagingrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
)

func (r *messagingRepositoryImpl) MarkMessagesAsRead(convID uint, userID uint, userType string) error {
	now := time.Now()
	return r.db.Model(&models.Message{}).
		Where("conversation_id = ? AND sender_id != ? AND sender_type != ? AND is_read = ?",
			convID, userID, userType, false).
		Updates(map[string]interface{}{
			"is_read": true,
			"read_at": now,
		}).Error
}
