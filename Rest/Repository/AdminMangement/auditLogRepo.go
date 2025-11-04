package adminmangement

import (
	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

type AuditLogRepository struct {
	db *gorm.DB
}

func NewAuditLogRepository(db *gorm.DB) *AuditLogRepository {
	return &AuditLogRepository{db: db}
}

// Create creates a new Audit log entry
func (r *AuditLogRepository) Create(log *models.AuditLog) error {
	return r.db.Create(log).Error
}

// GetByAdmin retrieves Audit logs for a specific admin
func (r *AuditLogRepository) GetByAdmin(adminID uint, page, limit int) ([]models.AuditLog, int64, error) {
	var logs []models.AuditLog
	var total int64

	query := r.db.Model(&models.AuditLog{}).
		Where("admin_id = ?", adminID).
		Preload("Admin")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&logs).Error

	return logs, total, err
}

// GetAll retrieves all Audit logs with pagination
func (r *AuditLogRepository) GetAll(page, limit int, action, entityType string) ([]models.AuditLog, int64, error) {
	var logs []models.AuditLog
	var total int64

	query := r.db.Model(&models.AuditLog{}).Preload("Admin")

	if action != "" {
		query = query.Where("action = ?", action)
	}
	if entityType != "" {
		query = query.Where("entity_type = ?", entityType)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&logs).Error

	return logs, total, err
}

// GetRecentAudit retrieves recent Audit logs
func (r *AuditLogRepository) GetRecentAudit(limit int) ([]models.AuditLog, error) {
	var logs []models.AuditLog
	err := r.db.Preload("Admin").Order("created_at DESC").Limit(limit).Find(&logs).Error
	return logs, err
}

// DeleteOldLogs deletes logs older than specified days
func (r *AuditLogRepository) DeleteOldLogs(days int) error {
	return r.db.Where("created_at < NOW() - INTERVAL '? days'", days).Delete(&models.AuditLog{}).Error
}
