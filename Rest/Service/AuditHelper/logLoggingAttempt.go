package audithelper

import (
	"fmt"

	"github.com/saadahmedbd/Treestore/constants"
)

// Log login attempt
func (h *AuditHelper) LogLoginAttempt(userID *uint, email, ip, userAgent string, success bool) {
	action := constants.ActionLogin
	severity := constants.SeverityInfo

	if !success {
		action = constants.ActionLoginFailed
		severity = constants.SeverityWarning
	}

	h.auditService.LogSecurityEvent(
		userID,
		"user",
		action,
		fmt.Sprintf("Login attempt for %s", email),
		ip,
		userAgent,

		severity,
	)
}
