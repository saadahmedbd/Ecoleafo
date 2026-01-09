package audithelper

import (
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	auditlogdto "github.com/saadahmedbd/Treestore/Rest/DTO/AuditLogDTO"
	"github.com/saadahmedbd/Treestore/constants"
)

// Log order status change
func (h *AuditHelper) LogOrderStatusChange(actorID uint, actorType, actorName string, order *models.Order, oldStatus, newStatus string) {
	h.auditService.Log(&auditlogdto.CreateAuditLogRequest{
		Action:      constants.ActionOrderStatusChange,
		Description: fmt.Sprintf("Order %s status changed from %s to %s", order.OrderNumber, oldStatus, newStatus),
		EntityType:  "order",
		EntityID:    &order.ID,
		EntityName:  order.OrderNumber,
		OldValues:   map[string]interface{}{"status": oldStatus},
		NewValues:   map[string]interface{}{"status": newStatus},
	}, actorID, actorType, actorName, "", "", "")
}
