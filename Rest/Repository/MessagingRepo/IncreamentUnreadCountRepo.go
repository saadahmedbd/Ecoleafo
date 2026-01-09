package messagingrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *messagingRepositoryImpl) IncrementUnreadCount(convID uint, participantNumber int) error {
	field := "participant1_unread_count"
	if participantNumber == 2 {
		field = "participant2_unread_count"
	}
	return r.db.Model(&models.Conversation{}).Where("id = ?", convID).
		UpdateColumn(field, gorm.Expr(field+" + ?", 1)).Error
}
