package auditlogrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
)

// Count by date range
func (r *auditLogRepository) CountByDateRange(startDate, endDate time.Time) (int64, error) {
	var count int64
	err := r.db.Model(&models.AuditLog{}).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Count(&count).Error
	return count, err
}
