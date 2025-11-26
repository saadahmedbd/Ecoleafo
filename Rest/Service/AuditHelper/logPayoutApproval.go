package audithelper

import (
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	"github.com/saadahmedbd/Treestore/constants"
)

// Log payout approval
func (h *AuditHelper) LogPayoutApproval(adminID uint, adminName string, payout *models.SellerPayout) {
	h.auditService.LogAdminAction(
		adminID,
		adminName,
		constants.ActionPayoutApprove,
		"payout",
		payout.ID,
		fmt.Sprintf("$%.2f", payout.Amount),
		fmt.Sprintf("Approved payout of $%.2f for seller #%d", payout.Amount, payout.SellerID),
	)
}
