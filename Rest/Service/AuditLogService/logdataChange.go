package auditlogservice

import (
	"encoding/json"

	models "github.com/saadahmedbd/Treestore/Models"
	"github.com/saadahmedbd/Treestore/constants"
)

// Quick helper: Log data change with diff
func (s *auditLogService) LogDataChange(actorID uint, actorType, action, entityType string, entityID uint, oldData, newData interface{}) error {
	oldJSON, _ := json.Marshal(oldData)
	newJSON, _ := json.Marshal(newData)

	var oldMap, newMap map[string]interface{}
	json.Unmarshal(oldJSON, &oldMap)
	json.Unmarshal(newJSON, &newMap)

	changes := s.calculateChanges(oldMap, newMap)
	changesJSON, _ := json.Marshal(changes)

	log := &models.AuditLog{
		ActorID:     &actorID,
		ActorType:   actorType,
		Action:      action,
		ActionGroup: s.getActionGroup(action),
		EntityType:  entityType,
		EntityID:    &entityID,
		OldValues:   string(oldJSON),
		NewValues:   string(newJSON),
		Changes:     string(changesJSON),
		Status:      "success",
		Severity:    constants.SeverityInfo,
		Category:    constants.CategoryAudit,
	}
	return s.auditRepo.Create(log)
}
