package auditmiddleware

import "strings"

func (m *AuditMiddleware) shouldSkip(path string) bool {
	skipPaths := []string{
		"/api/health",
		"/api/notifications/unread-count",
		"/api/admin/logs",
	}
	for _, p := range skipPaths {
		if strings.HasPrefix(path, p) {
			return true
		}
	}
	return false
}
