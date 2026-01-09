package auditlogservice

import (
	"encoding/json"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	auditlogdto "github.com/saadahmedbd/Treestore/Rest/DTO/AuditLogDTO"
)

// Helper: Convert to response
func (s *auditLogService) toResponse(log *models.AuditLog) *auditlogdto.AuditLogResponse {
	resp := &auditlogdto.AuditLogResponse{
		ID:          log.ID,
		ActorID:     log.ActorID,
		ActorType:   log.ActorType,
		ActorName:   log.ActorName,
		ActorEmail:  log.ActorEmail,
		Action:      log.Action,
		ActionGroup: log.ActionGroup,
		ActionLabel: s.getActionLabel(log.Action),
		Description: log.Description,
		EntityType:  log.EntityType,
		EntityID:    log.EntityID,
		EntityName:  log.EntityName,
		IPAddress:   log.IPAddress,
		UserAgent:   log.UserAgent,
		Status:      log.Status,
		Severity:    log.Severity,
		Category:    log.Category,
		CreatedAt:   log.CreatedAt.Format(time.RFC3339),
		TimeAgo:     s.timeAgo(log.CreatedAt),
	}

	// Parse JSON fields
	if log.OldValues != "" {
		json.Unmarshal([]byte(log.OldValues), &resp.OldValues)
	}
	if log.NewValues != "" {
		json.Unmarshal([]byte(log.NewValues), &resp.NewValues)
	}
	if log.Changes != "" {
		json.Unmarshal([]byte(log.Changes), &resp.Changes)
	}

	return resp
}
