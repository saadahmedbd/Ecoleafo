package auditlogrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	auditlogdto "github.com/saadahmedbd/Treestore/Rest/DTO/AuditLogDTO"
)

// Get all audit logs with filters
func (r *auditLogRepository) GetAll(query auditlogdto.AuditLogQuery) ([]models.AuditLog, int64, error) {
	var logs []models.AuditLog
	var total int64

	db := r.db.Model(&models.AuditLog{})

	// Apply filters
	if query.ActorID != nil {
		db = db.Where("actor_id = ?", *query.ActorID)
	}
	if query.ActorType != "" {
		db = db.Where("actor_type = ?", query.ActorType)
	}
	if query.Action != "" {
		db = db.Where("action = ?", query.Action)
	}
	if query.ActionGroup != "" {
		db = db.Where("action_group = ?", query.ActionGroup)
	}
	if query.EntityType != "" {
		db = db.Where("entity_type = ?", query.EntityType)
	}
	if query.EntityID != nil {
		db = db.Where("entity_id = ?", *query.EntityID)
	}
	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}
	if query.Severity != "" {
		db = db.Where("severity = ?", query.Severity)
	}
	if query.Category != "" {
		db = db.Where("category = ?", query.Category)
	}
	if query.IPAddress != "" {
		db = db.Where("ip_address = ?", query.IPAddress)
	}
	if query.StartDate != "" {
		startDate, _ := time.Parse("2006-01-02", query.StartDate)
		db = db.Where("created_at >= ?", startDate)
	}
	if query.EndDate != "" {
		endDate, _ := time.Parse("2006-01-02", query.EndDate)
		endDate = endDate.Add(24 * time.Hour) // Include full day
		db = db.Where("created_at < ?", endDate)
	}
	if query.Search != "" {
		search := "%" + query.Search + "%"
		db = db.Where("description ILIKE ? OR actor_name ILIKE ? OR entity_name ILIKE ?",
			search, search, search)
	}

	// Count total
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	page := query.Page
	limit := query.Limit
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}
	offset := (page - 1) * limit

	// Get results
	err := db.Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&logs).Error

	return logs, total, err
}
