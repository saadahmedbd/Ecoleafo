package audithelper

import (
	"fmt"

	auditlogdto "github.com/saadahmedbd/Treestore/Rest/DTO/AuditLogDTO"
	"github.com/saadahmedbd/Treestore/constants"
)

// Log settings update
func (h *AuditHelper) LogSettingsUpdate(adminID uint, adminName string, settingKey string, oldValue, newValue interface{}) {
	h.auditService.Log(&auditlogdto.CreateAuditLogRequest{
		Action:      constants.ActionSettingsUpdate,
		Description: fmt.Sprintf("Updated setting: %s", settingKey),
		EntityType:  "setting",
		EntityName:  settingKey,
		OldValues:   map[string]interface{}{settingKey: oldValue},
		NewValues:   map[string]interface{}{settingKey: newValue},
		Category:    constants.CategorySystem,
	}, adminID, "admin", adminName, "", "", "")
}
