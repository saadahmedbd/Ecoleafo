package auditlogrepo

import models "github.com/saadahmedbd/Treestore/Models"

// Get logs by action
func (r *auditLogRepository) GetByAction(action string, page, limit int) ([]models.AuditLog, int64, error) {
	var logs []models.AuditLog
	var total int64

	offset := (page - 1) * limit

	db := r.db.Model(&models.AuditLog{}).Where("action = ?", action)

	db.Count(&total)
	err := db.Order("created_at DESC").Offset(offset).Limit(limit).Find(&logs).Error

	return logs, total, err
}
