package auditlogservice

import auditlogdto "github.com/saadahmedbd/Treestore/Rest/DTO/AuditLogDTO"

// Get actor logs
func (s *auditLogService) GetActorLogs(actorID uint, actorType string, page, limit int) ([]auditlogdto.AuditLogResponse, int64, error) {
	logs, total, err := s.auditRepo.GetByActor(actorID, actorType, page, limit)
	if err != nil {
		return nil, 0, err
	}
	return s.toResponseList(logs), total, nil
}
