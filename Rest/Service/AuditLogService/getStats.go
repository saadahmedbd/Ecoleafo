package auditlogservice

import (
	"time"

	auditlogdto "github.com/saadahmedbd/Treestore/Rest/DTO/AuditLogDTO"
)

// Get statistics
func (s *auditLogService) GetStats() (*auditlogdto.AuditLogStatsResponse, error) {
	startDate := time.Now().AddDate(0, 0, -30)
	endDate := time.Now()

	stats, err := s.auditRepo.GetStats(startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Get recent activity
	recentLogs, _ := s.auditRepo.GetRecentActivity(10)
	recentActivity := s.toResponseList(recentLogs)

	// Get top actors
	topActors, _ := s.auditRepo.GetTopActors(startDate, endDate, 10)

	return &auditlogdto.AuditLogStatsResponse{
		TotalLogs:      int(stats["total_logs"].(int64)),
		TodayLogs:      int(stats["today_logs"].(int64)),
		ThisWeekLogs:   int(stats["this_week_logs"].(int64)),
		SuccessCount:   int(stats["success_count"].(int64)),
		FailedCount:    int(stats["failed_count"].(int64)),
		ByActionGroup:  stats["by_action_group"].(map[string]int),
		BySeverity:     stats["by_severity"].(map[string]int),
		ByActorType:    stats["by_actor_type"].(map[string]int),
		RecentActivity: recentActivity,
		TopActors:      topActors,
		SecurityAlerts: int(stats["security_alerts"].(int64)),
	}, nil
}
