package auditlogrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *auditLogRepository) Create(log *models.AuditLog) error {
	return r.db.Create(log).Error
}
