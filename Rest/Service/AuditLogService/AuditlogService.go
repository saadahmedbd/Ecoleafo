package auditlogservice

import (
	auditlogdto "github.com/saadahmedbd/Treestore/Rest/DTO/AuditLogDTO"
	auditlogrepo "github.com/saadahmedbd/Treestore/Rest/Repository/AuditLogRepo"
)

type AuditLogService interface {
	// Logging methods
	Log(req *auditlogdto.CreateAuditLogRequest, actorID uint, actorType, actorName, actorEmail, ip, userAgent string) error
	// LogFromContext(ctx context.Context, req *auditlogdto.CreateAuditLogRequest) error

	// Quick logging helpers
	LogAdminAction(adminID uint, adminName, action, entityType string, entityID uint, entityName, description string) error
	// LogSellerAction(sellerID uint, sellerName, action, entityType string, entityID uint, entityName, description string) error
	LogSecurityEvent(actorID *uint, actorType, action, description, ip, userAgent string, severity string) error
	LogDataChange(actorID uint, actorType, action, entityType string, entityID uint, oldData, newData interface{}) error

	// Query methods
	GetLogs(query auditlogdto.AuditLogQuery) ([]auditlogdto.AuditLogResponse, int64, error)
	GetLogByID(id uint) (*auditlogdto.AuditLogResponse, error)
	GetActorLogs(actorID uint, actorType string, page, limit int) ([]auditlogdto.AuditLogResponse, int64, error)
	GetEntityHistory(entityType string, entityID uint, page, limit int) ([]auditlogdto.AuditLogResponse, int64, error)
	GetSecurityLogs(page, limit int) ([]auditlogdto.AuditLogResponse, int64, error)
	GetRecentActivity(limit int) ([]auditlogdto.AuditLogResponse, error)

	// Statistics
	GetStats() (*auditlogdto.AuditLogStatsResponse, error)
	GetActivityTimeline(startDate, endDate string, groupBy string) ([]auditlogdto.ActivityTimelineItem, error)

	// Export
	ExportLogs(req auditlogdto.ExportAuditLogsRequest) ([]byte, string, error)

	// Maintenance
	CleanupOldLogs(days int) error
}

type auditLogService struct {
	auditRepo auditlogrepo.AuditLogRepository
}

func NewAuditLogService(auditRepo auditlogrepo.AuditLogRepository) AuditLogService {
	return &auditLogService{
		auditRepo: auditRepo,
	}
}
