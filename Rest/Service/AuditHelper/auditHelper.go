package audithelper

import auditlogservice "github.com/saadahmedbd/Treestore/Rest/Service/AuditLogService"

type AuditHelper struct {
	auditService auditlogservice.AuditLogService
}

func NewAuditHelper(auditService auditlogservice.AuditLogService) *AuditHelper {
	return &AuditHelper{auditService: auditService}
}
