package auditloghandler

import auditlogservice "github.com/saadahmedbd/Treestore/Rest/Service/AuditLogService"

type AuditLogHandler struct {
	auditService auditlogservice.AuditLogService
}

func NewAuditLogHandler(auditService auditlogservice.AuditLogService) *AuditLogHandler {
	return &AuditLogHandler{auditService: auditService}
}
