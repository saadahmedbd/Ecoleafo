package auditlogservice

// Cleanup old logs
func (s *auditLogService) CleanupOldLogs(days int) error {
	if days < 30 {
		days = 30 // Minimum retention
	}
	return s.auditRepo.DeleteOldLogs(days)
}
