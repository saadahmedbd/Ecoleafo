package auditmiddleware

import auditlogservice "github.com/saadahmedbd/Treestore/Rest/Service/AuditLogService"

type AuditMiddleware struct {
	auditService auditlogservice.AuditLogService
}

func NewAuditMiddleware(auditService auditlogservice.AuditLogService) *AuditMiddleware {
	return &AuditMiddleware{auditService: auditService}
}
