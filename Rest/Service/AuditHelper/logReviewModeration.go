package audithelper

import (
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	"github.com/saadahmedbd/Treestore/constants"
)

// Log review moderation
func (h *AuditHelper) LogReviewModeration(adminID uint, adminName string, review *models.Review, action string, reason string) {
	actionType := constants.ActionReviewApprove
	if action == "reject" {
		actionType = constants.ActionReviewReject
	} else if action == "delete" {
		actionType = constants.ActionReviewDelete
	}

	desc := fmt.Sprintf("Review #%d %s", review.ID, action)
	if reason != "" {
		desc += fmt.Sprintf(". Reason: %s", reason)
	}

	h.auditService.LogAdminAction(
		adminID,
		adminName,
		actionType,
		"review",
		review.ID,
		fmt.Sprintf("Review for product #%d", review.ProductID),
		desc,
	)
}
