package audithelper

import (
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	"github.com/saadahmedbd/Treestore/constants"
)

// Log seller approval
func (h *AuditHelper) LogSellerApproval(adminID uint, adminName string, seller *models.User) {
	h.auditService.LogAdminAction(
		adminID,
		adminName,
		constants.ActionSellerApprove,
		"seller",
		seller.ID,
		seller.StoreName,
		fmt.Sprintf("Approved seller: %s", seller.StoreName),
	)
}
