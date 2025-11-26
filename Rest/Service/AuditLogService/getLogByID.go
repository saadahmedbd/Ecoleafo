package auditlogservice

import auditlogdto "github.com/saadahmedbd/Treestore/Rest/DTO/AuditLogDTO"

// Get log by ID
func (s *auditLogService) GetLogByID(id uint) (*auditlogdto.AuditLogResponse, error) {
	log, err := s.auditRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return s.toResponse(log), nil
}
