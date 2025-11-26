package auditlogservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	"github.com/saadahmedbd/Treestore/constants"
)

// Quick helper: Log security event
func (s *auditLogService) LogSecurityEvent(actorID *uint, actorType, action, description, ip, userAgent string, severity string) error {
	log := &models.AuditLog{
		ActorID:     actorID,
		ActorType:   actorType,
		Action:      action,
		ActionGroup: constants.GroupSecurity,
		Description: description,
		IPAddress:   ip,
		UserAgent:   userAgent,
		Status:      "success",
		Severity:    severity,
		Category:    constants.CategorySecurity,
	}
	return s.auditRepo.Create(log)
}
