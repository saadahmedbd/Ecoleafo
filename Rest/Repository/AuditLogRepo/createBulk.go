package auditlogrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *auditLogRepository) CreateBulk(logs []models.AuditLog) error {
	if len(logs) == 0 {
		return nil
	}
	return r.db.CreateInBatches(logs, 100).Error
}
