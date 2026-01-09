package auditlogrepo

import models "github.com/saadahmedbd/Treestore/Models"

// Get recent activity
func (r *auditLogRepository) GetRecentActivity(limit int) ([]models.AuditLog, error) {
	var logs []models.AuditLog
	err := r.db.Order("created_at DESC").Limit(limit).Find(&logs).Error
	return logs, err
}
