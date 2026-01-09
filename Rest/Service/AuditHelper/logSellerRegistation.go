package audithelper

import (
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	"github.com/saadahmedbd/Treestore/constants"
)

// Log seller rejection
func (h *AuditHelper) LogSellerRejection(adminID uint, adminName string, seller *models.User, reason string) {
	h.auditService.LogAdminAction(
		adminID,
		adminName,
		constants.ActionSellerReject,
		"seller",
		seller.ID,
		seller.StoreName,
		fmt.Sprintf("Rejected seller: %s. Reason: %s", seller.StoreName, reason),
	)
}
