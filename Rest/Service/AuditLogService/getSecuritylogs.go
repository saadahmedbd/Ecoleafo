package auditlogservice

import auditlogdto "github.com/saadahmedbd/Treestore/Rest/DTO/AuditLogDTO"

// Get security logs
func (s *auditLogService) GetSecurityLogs(page, limit int) ([]auditlogdto.AuditLogResponse, int64, error) {
	logs, total, err := s.auditRepo.GetSecurityLogs(page, limit)
	if err != nil {
		return nil, 0, err
	}
	return s.toResponseList(logs), total, nil
}
