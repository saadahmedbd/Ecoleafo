package auditlogrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	"github.com/saadahmedbd/Treestore/constants"
)

// Get statistics
func (r *auditLogRepository) GetStats(startDate, endDate time.Time) (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// Total logs
	var totalLogs int64
	r.db.Model(&models.AuditLog{}).Where("created_at BETWEEN ? AND ?", startDate, endDate).Count(&totalLogs)
	stats["total_logs"] = totalLogs

	// Today's logs
	today := time.Now().Truncate(24 * time.Hour)
	var todayLogs int64
	r.db.Model(&models.AuditLog{}).Where("created_at >= ?", today).Count(&todayLogs)
	stats["today_logs"] = todayLogs

	// This week's logs
	weekAgo := time.Now().AddDate(0, 0, -7)
	var weekLogs int64
	r.db.Model(&models.AuditLog{}).Where("created_at >= ?", weekAgo).Count(&weekLogs)
	stats["this_week_logs"] = weekLogs

	// Success vs Failed
	var successCount, failedCount int64
	r.db.Model(&models.AuditLog{}).Where("status = ?", "success").Count(&successCount)
	r.db.Model(&models.AuditLog{}).Where("status = ?", "failed").Count(&failedCount)
	stats["success_count"] = successCount
	stats["failed_count"] = failedCount

	// By action group
	var actionGroupStats []struct {
		ActionGroup string
		Count       int64
	}
	r.db.Model(&models.AuditLog{}).
		Select("action_group, COUNT(*) as count").
		Group("action_group").
		Scan(&actionGroupStats)

	byActionGroup := make(map[string]int)
	for _, s := range actionGroupStats {
		byActionGroup[s.ActionGroup] = int(s.Count)
	}
	stats["by_action_group"] = byActionGroup

	// By severity
	var severityStats []struct {
		Severity string
		Count    int64
	}
	r.db.Model(&models.AuditLog{}).
		Select("severity, COUNT(*) as count").
		Group("severity").
		Scan(&severityStats)

	bySeverity := make(map[string]int)
	for _, s := range severityStats {
		bySeverity[s.Severity] = int(s.Count)
	}
	stats["by_severity"] = bySeverity

	// By actor type
	var actorStats []struct {
		ActorType string
		Count     int64
	}
	r.db.Model(&models.AuditLog{}).
		Select("actor_type, COUNT(*) as count").
		Group("actor_type").
		Scan(&actorStats)

	byActorType := make(map[string]int)
	for _, s := range actorStats {
		byActorType[s.ActorType] = int(s.Count)
	}
	stats["by_actor_type"] = byActorType

	// Security alerts count
	var securityAlerts int64
	r.db.Model(&models.AuditLog{}).
		Where("severity = ? OR category = ?", constants.SeverityCritical, constants.CategorySecurity).
		Count(&securityAlerts)
	stats["security_alerts"] = securityAlerts

	return stats, nil
}
