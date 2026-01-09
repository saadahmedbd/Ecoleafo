package auditlogservice

import "github.com/saadahmedbd/Treestore/constants"

// Helper: Get action group from action
func (s *auditLogService) getActionGroup(action string) string {
	groups := map[string]string{
		constants.ActionLogin:            constants.GroupAuth,
		constants.ActionLogout:           constants.GroupAuth,
		constants.ActionLoginFailed:      constants.GroupAuth,
		constants.ActionPasswordChange:   constants.GroupAuth,
		constants.ActionPasswordReset:    constants.GroupAuth,
		constants.ActionUserCreate:       constants.GroupUser,
		constants.ActionUserUpdate:       constants.GroupUser,
		constants.ActionUserDelete:       constants.GroupUser,
		constants.ActionUserActivate:     constants.GroupUser,
		constants.ActionUserDeactivate:   constants.GroupUser,
		constants.ActionSellerApprove:    constants.GroupSeller,
		constants.ActionSellerReject:     constants.GroupSeller,
		constants.ActionSellerSuspend:    constants.GroupSeller,
		constants.ActionProductCreate:    constants.GroupProduct,
		constants.ActionProductUpdate:    constants.GroupProduct,
		constants.ActionProductApprove:   constants.GroupProduct,
		constants.ActionProductReject:    constants.GroupProduct,
		constants.ActionOrderCreate:      constants.GroupOrder,
		constants.ActionOrderUpdate:      constants.GroupOrder,
		constants.ActionOrderCancel:      constants.GroupOrder,
		constants.ActionReviewApprove:    constants.GroupReview,
		constants.ActionReviewReject:     constants.GroupReview,
		constants.ActionPayoutApprove:    constants.GroupPayment,
		constants.ActionPayoutReject:     constants.GroupPayment,
		constants.ActionAdminInvite:      constants.GroupAdmin,
		constants.ActionPermissionUpdate: constants.GroupAdmin,
		constants.ActionSettingsUpdate:   constants.GroupSystem,
	}

	if group, ok := groups[action]; ok {
		return group
	}
	return action
}
