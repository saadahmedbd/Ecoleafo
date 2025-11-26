package auditlogrepo

import (
	"time"

	auditlogdto "github.com/saadahmedbd/Treestore/Rest/DTO/AuditLogDTO"
)

// Get top actors
func (r *auditLogRepository) GetTopActors(startDate, endDate time.Time, limit int) ([]auditlogdto.TopActorItem, error) {
	var items []auditlogdto.TopActorItem

	query := `
		SELECT 
			actor_id,
			actor_name,
			actor_type,
			COUNT(*) as action_count
		FROM audit_logs
		WHERE created_at BETWEEN ? AND ?
		AND actor_id IS NOT NULL
		GROUP BY actor_id, actor_name, actor_type
		ORDER BY action_count DESC
		LIMIT ?
	`

	err := r.db.Raw(query, startDate, endDate, limit).Scan(&items).Error
	return items, err
}
