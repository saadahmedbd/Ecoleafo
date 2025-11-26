package auditlogrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	auditlogdto "github.com/saadahmedbd/Treestore/Rest/DTO/AuditLogDTO"
	"gorm.io/gorm"
)

type AuditLogRepository interface {
	Create(log *models.AuditLog) error
	CreateBulk(logs []models.AuditLog) error
	GetByID(id uint) (*models.AuditLog, error)
	GetAll(query auditlogdto.AuditLogQuery) ([]models.AuditLog, int64, error)

	// Specialized Queries
	GetByActor(actorID uint, actorType string, page, limit int) ([]models.AuditLog, int64, error)
	GetByEntity(entityType string, entityID uint, page, limit int) ([]models.AuditLog, int64, error)
	GetByAction(action string, page, limit int) ([]models.AuditLog, int64, error)
	GetSecurityLogs(page, limit int) ([]models.AuditLog, int64, error)
	GetRecentActivity(limit int) ([]models.AuditLog, error)

	// Statistics
	GetStats(startDate, endDate time.Time) (map[string]interface{}, error)
	GetActivityTimeline(startDate, endDate time.Time, groupBy string) ([]auditlogdto.ActivityTimelineItem, error)
	GetTopActors(startDate, endDate time.Time, limit int) ([]auditlogdto.TopActorItem, error)

	// Maintenance
	DeleteOldLogs(days int) error
	CountByDateRange(startDate, endDate time.Time) (int64, error)
}

type auditLogRepository struct {
	db *gorm.DB
}

func NewAuditLogRepository(db *gorm.DB) AuditLogRepository {
	return &auditLogRepository{
		db: db,
	}
}
