package auditlogrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	"github.com/saadahmedbd/Treestore/constants"
)

// Get security logs (critical severity or security category)
func (r *auditLogRepository) GetSecurityLogs(page, limit int) ([]models.AuditLog, int64, error) {
	var logs []models.AuditLog
	var total int64

	offset := (page - 1) * limit

	db := r.db.Model(&models.AuditLog{}).
		Where("category = ? OR severity = ? OR action IN ?",
			constants.CategorySecurity, constants.SeverityCritical,
			[]string{constants.ActionLoginFailed, constants.ActionPasswordChange, constants.ActionPasswordReset})

	db.Count(&total)
	err := db.Order("created_at DESC").Offset(offset).Limit(limit).Find(&logs).Error

	return logs, total, err
}
