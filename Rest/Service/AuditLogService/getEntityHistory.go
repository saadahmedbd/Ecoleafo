package auditlogservice

import auditlogdto "github.com/saadahmedbd/Treestore/Rest/DTO/AuditLogDTO"

// Get entity history
func (s *auditLogService) GetEntityHistory(entityType string, entityID uint, page, limit int) ([]auditlogdto.AuditLogResponse, int64, error) {
	logs, total, err := s.auditRepo.GetByEntity(entityType, entityID, page, limit)
	if err != nil {
		return nil, 0, err
	}
	return s.toResponseList(logs), total, nil
}
