package messagingrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
)

func (r *messagingRepositoryImpl) GetMessagesByConversationID(convID uint, since time.Time, limit int) ([]models.Message, error) {
	var msgs []models.Message
	query := r.db.Where("conversation_id = ?", convID)

	if !since.IsZero() {
		query = query.Where("created_at > ?", since)
	}

	if limit > 0 {
		query = query.Limit(limit)
	} else {
		query = query.Limit(50)
	}

	err := query.Order("created_at ASC").Find(&msgs).Error
	return msgs, err
}
