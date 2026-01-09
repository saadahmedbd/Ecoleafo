package auditlogservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	"github.com/saadahmedbd/Treestore/constants"
)

// Quick helper: Log admin action
func (s *auditLogService) LogAdminAction(adminID uint, adminName, action, entityType string, entityID uint, entityName, description string) error {
	log := &models.AuditLog{
		ActorID:     &adminID,
		ActorType:   "admin",
		ActorName:   adminName,
		Action:      action,
		ActionGroup: s.getActionGroup(action),
		Description: description,
		EntityType:  entityType,
		EntityID:    &entityID,
		EntityName:  entityName,
		Status:      "success",
		Severity:    constants.SeverityInfo,
		Category:    constants.CategoryBusiness,
	}
	return s.auditRepo.Create(log)
}
