package auditlogservice

import (
	"encoding/json"

	models "github.com/saadahmedbd/Treestore/Models"
	auditlogdto "github.com/saadahmedbd/Treestore/Rest/DTO/AuditLogDTO"
)

// Main logging method
func (s *auditLogService) Log(req *auditlogdto.CreateAuditLogRequest, actorID uint, actorType, actorName, actorEmail, ip, userAgent string) error {
	// Marshal JSON fields
	oldValuesJSON := ""
	if req.OldValues != nil {
		bytes, _ := json.Marshal(req.OldValues)
		oldValuesJSON = string(bytes)
	}

	newValuesJSON := ""
	if req.NewValues != nil {
		bytes, _ := json.Marshal(req.NewValues)
		newValuesJSON = string(bytes)
	}

	metadataJSON := ""
	if req.Metadata != nil {
		bytes, _ := json.Marshal(req.Metadata)
		metadataJSON = string(bytes)
	}

	// Calculate changes
	changesJSON := ""
	if req.OldValues != nil && req.NewValues != nil {
		changes := s.calculateChanges(req.OldValues, req.NewValues)
		bytes, _ := json.Marshal(changes)
		changesJSON = string(bytes)
	}

	// Set defaults
	if req.Status == "" {
		req.Status = "success"
	}
	if req.Severity == "" {
		req.Severity = req.Severity
	}
	if req.Category == "" {
		req.Category = req.Category
	}

	// Create audit log
	log := &models.AuditLog{
		ActorID:     &actorID,
		ActorType:   actorType,
		ActorName:   actorName,
		ActorEmail:  actorEmail,
		Action:      req.Action,
		ActionGroup: s.getActionGroup(req.Action),
		Description: req.Description,
		EntityType:  req.EntityType,
		EntityID:    req.EntityID,
		EntityName:  req.EntityName,
		OldValues:   oldValuesJSON,
		NewValues:   newValuesJSON,
		Changes:     changesJSON,
		IPAddress:   ip,
		UserAgent:   userAgent,
		Status:      req.Status,
		Severity:    req.Severity,
		Category:    req.Category,
		Metadata:    metadataJSON,
	}

	return s.auditRepo.Create(log)
}
