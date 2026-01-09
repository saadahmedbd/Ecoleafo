package auditlogservice

import "github.com/saadahmedbd/Treestore/constants"

// Helper: Get human readable action label
func (s *auditLogService) getActionLabel(action string) string {
	labels := map[string]string{
		constants.ActionLogin:            "User Login",
		constants.ActionLogout:           "User Logout",
		constants.ActionLoginFailed:      "Failed Login Attempt",
		constants.ActionPasswordChange:   "Password Changed",
		constants.ActionPasswordReset:    "Password Reset",
		constants.ActionUserCreate:       "User Created",
		constants.ActionUserUpdate:       "User Updated",
		constants.ActionUserDelete:       "User Deleted",
		constants.ActionUserActivate:     "User Activated",
		constants.ActionUserDeactivate:   "User Deactivated",
		constants.ActionUserSuspend:      "User Suspended",
		constants.ActionSellerApprove:    "Seller Approved",
		constants.ActionSellerReject:     "Seller Rejected",
		constants.ActionSellerSuspend:    "Seller Suspended",
		constants.ActionSellerReactivate: "Seller Reactivated",
		constants.ActionProductCreate:    "Product Created",
		constants.ActionProductUpdate:    "Product Updated",
		constants.ActionProductDelete:    "Product Deleted",
		constants.ActionProductApprove:   "Product Approved",
		constants.ActionProductReject:    "Product Rejected",
		constants.ActionOrderCreate:      "Order Created",
		constants.ActionOrderUpdate:      "Order Updated",
		constants.ActionOrderCancel:      "Order Cancelled",
		constants.ActionOrderRefund:      "Order Refunded",
		constants.ActionReviewApprove:    "Review Approved",
		constants.ActionReviewReject:     "Review Rejected",
		constants.ActionReviewDelete:     "Review Deleted",
		constants.ActionPayoutRequest:    "Payout Requested",
		constants.ActionPayoutApprove:    "Payout Approved",
		constants.ActionPayoutReject:     "Payout Rejected",
		constants.ActionAdminInvite:      "Admin Invited",
		constants.ActionAdminCreate:      "Admin Created",
		constants.ActionPermissionUpdate: "Permissions Updated",
		constants.ActionSettingsUpdate:   "Settings Updated",
		constants.ActionDataExport:       "Data Exported",
	}

	if label, ok := labels[action]; ok {
		return label
	}
	return action
}
