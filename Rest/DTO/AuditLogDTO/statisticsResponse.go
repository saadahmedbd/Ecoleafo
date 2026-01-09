package auditlogdto

// Statistics response
type AuditLogStatsResponse struct {
	TotalLogs      int                `json:"total_logs"`
	TodayLogs      int                `json:"today_logs"`
	ThisWeekLogs   int                `json:"this_week_logs"`
	SuccessCount   int                `json:"success_count"`
	FailedCount    int                `json:"failed_count"`
	ByActionGroup  map[string]int     `json:"by_action_group"`
	BySeverity     map[string]int     `json:"by_severity"`
	ByActorType    map[string]int     `json:"by_actor_type"`
	RecentActivity []AuditLogResponse `json:"recent_activity"`
	TopActors      []TopActorItem     `json:"top_actors"`
	SecurityAlerts int                `json:"security_alerts"`
}
