package audithelper

import (
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	"github.com/saadahmedbd/Treestore/constants"
)

// Log product approval
func (h *AuditHelper) LogProductApproval(adminID uint, adminName string, product *models.Product) {
	h.auditService.LogAdminAction(
		adminID,
		adminName,
		constants.ActionProductApprove,
		"product",
		product.ID,
		product.Name,
		fmt.Sprintf("Approved product: %s", product.Name),
	)
}
