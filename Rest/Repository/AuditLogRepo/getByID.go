package auditlogrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *auditLogRepository) GetByID(id uint) (*models.AuditLog, error) {
	var log models.AuditLog
	err := r.db.First(&log, id).Error
	return &log, err
}
