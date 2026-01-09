package messagingrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *messagingRepositoryImpl) GetConversationByParticipants(p1ID uint, p1Type string, p2ID uint, p2Type string, contextType string, contextID *uint) (*models.Conversation, error) {
	var conv models.Conversation
	query := r.db.Where(
		"((participant1_id = ? AND participant1_type = ? AND participant2_id = ? AND participant2_type = ?) OR "+
			"(participant1_id = ? AND participant1_type = ? AND participant2_id = ? AND participant2_type = ?))",
		p1ID, p1Type, p2ID, p2Type,
		p2ID, p2Type, p1ID, p1Type,
	)

	if contextType != "" {
		query = query.Where("context_type = ?", contextType)
	}
	if contextID != nil {
		query = query.Where("context_id = ?", *contextID)
	}

	err := query.Preload("Product").Preload("Order").First(&conv).Error
	return &conv, err
}
