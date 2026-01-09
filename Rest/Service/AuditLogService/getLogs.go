package auditlogservice

import auditlogdto "github.com/saadahmedbd/Treestore/Rest/DTO/AuditLogDTO"

// Get logs with filters
func (s *auditLogService) GetLogs(query auditlogdto.AuditLogQuery) ([]auditlogdto.AuditLogResponse, int64, error) {
	if query.Page < 1 {
		query.Page = 1
	}
	if query.Limit < 1 || query.Limit > 100 {
		query.Limit = 20
	}

	logs, total, err := s.auditRepo.GetAll(query)
	if err != nil {
		return nil, 0, err
	}

	return s.toResponseList(logs), total, nil
}
