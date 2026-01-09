package audithelper

import (
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	"github.com/saadahmedbd/Treestore/constants"
)

// Log product rejection
func (h *AuditHelper) LogProductRejection(adminID uint, adminName string, product *models.Product, reason string) {
	h.auditService.LogAdminAction(
		adminID,
		adminName,
		constants.ActionProductReject,
		"product",
		product.ID,
		product.Name,
		fmt.Sprintf("Rejected product: %s. Reason: %s", product.Name, reason),
	)
}
