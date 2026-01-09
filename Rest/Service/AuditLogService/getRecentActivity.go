package auditlogservice

import auditlogdto "github.com/saadahmedbd/Treestore/Rest/DTO/AuditLogDTO"

// Get recent activity
func (s *auditLogService) GetRecentActivity(limit int) ([]auditlogdto.AuditLogResponse, error) {
	if limit < 1 || limit > 50 {
		limit = 20
	}
	logs, err := s.auditRepo.GetRecentActivity(limit)
	if err != nil {
		return nil, err
	}
	return s.toResponseList(logs), nil
}
