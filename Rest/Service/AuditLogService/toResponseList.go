package auditlogservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	auditlogdto "github.com/saadahmedbd/Treestore/Rest/DTO/AuditLogDTO"
)

// Helper: Convert list to responses
func (s *auditLogService) toResponseList(logs []models.AuditLog) []auditlogdto.AuditLogResponse {
	var responses []auditlogdto.AuditLogResponse
	for _, log := range logs {
		responses = append(responses, *s.toResponse(&log))
	}
	return responses
}
