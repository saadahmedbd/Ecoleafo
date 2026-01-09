package auditlogrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
)

// Delete old logs
func (r *auditLogRepository) DeleteOldLogs(days int) error {
	cutoff := time.Now().AddDate(0, 0, -days)
	return r.db.Where("created_at < ?", cutoff).Delete(&models.AuditLog{}).Error
}
