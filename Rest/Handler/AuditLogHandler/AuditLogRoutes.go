package auditloghandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *AuditLogHandler) RegisterAuditLog(mux *http.ServeMux) {
	// get all logs with filter

	mux.Handle("GET /api/admin/logs",
		middleware.Chain(
			http.HandlerFunc(h.GetLogs),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
			middleware.Requirerole([]string{"admin", "super_admin"}),
		),
	)
	//get log by id
	mux.Handle("GET /api/admin/logs/detail",
		middleware.Chain(
			http.HandlerFunc(h.GetLogByID),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
			middleware.Requirerole([]string{"admin", "super_admin"}),
		),
	)
	mux.Handle("GET /api/admin/logs/actor",
		middleware.Chain(
			http.HandlerFunc(h.GetActorLogs),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
			middleware.Requirerole([]string{"admin", "super_admin"}),
		),
	)
	//get entity history
	mux.Handle("GET /api/admin/logs/entity",
		middleware.Chain(
			http.HandlerFunc(h.GetEntityHistory),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
			middleware.Requirerole([]string{"admin", "super_admin"}),
		),
	)
	mux.Handle("GET /api/admin/logs/security",
		middleware.Chain(
			http.HandlerFunc(h.GetSecurityLogs),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
			middleware.Requirerole([]string{"admin", "super_admin"}),
		),
	)
	mux.Handle("GET /api/admin/logs/recent",
		middleware.Chain(
			http.HandlerFunc(h.GetRecentActivity),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
			middleware.Requirerole([]string{"admin", "super_admin"}),
		),
	)
	mux.Handle("GET /api/admin/logs/stats",
		middleware.Chain(
			http.HandlerFunc(h.GetStats),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
			middleware.Requirerole([]string{"admin", "super_admin"}),
		),
	)
	mux.Handle("GET /api/admin/logs/timeline",
		middleware.Chain(
			http.HandlerFunc(h.GetActivityTimeline),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
			middleware.Requirerole([]string{"admin", "super_admin"}),
		),
	)
	mux.Handle("POST /api/admin/logs/export",
		middleware.Chain(
			http.HandlerFunc(h.ExportLogs),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
			middleware.Requirerole([]string{"admin", "super_admin"}),
		),
	)
	mux.Handle("GET /api/admin/logs/cleanup",
		middleware.Chain(
			http.HandlerFunc(h.CleanupOldLogs),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
			middleware.Requirerole([]string{"admin", "super_admin"}),
		),
	)

}
