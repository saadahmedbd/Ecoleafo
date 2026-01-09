package auditmiddleware

import (
	"fmt"
	"net/http"

	auditlogdto "github.com/saadahmedbd/Treestore/Rest/DTO/AuditLogDTO"
)

// Middleware to auto-log requests
func (m *AuditMiddleware) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip logging for certain paths
		if m.shouldSkip(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		// Get user info from context
		actorID, _ := r.Context().Value("user_id").(uint)
		actorType, _ := r.Context().Value("role").(string)
		actorName, _ := r.Context().Value("user_name").(string)
		actorEmail, _ := r.Context().Value("user_email").(string)

		// Get request info
		ip := m.getClientIP(r)
		userAgent := r.UserAgent()

		// Determine action from request
		action := m.determineAction(r.Method, r.URL.Path)

		// Log the request
		m.auditService.Log(&auditlogdto.CreateAuditLogRequest{
			Action:      action,
			Description: fmt.Sprintf("%s %s", r.Method, r.URL.Path),
		}, actorID, actorType, actorName, actorEmail, ip, userAgent)

		next.ServeHTTP(w, r)
	})
}
