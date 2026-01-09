package auditlogrepo

import (
	"fmt"
	"time"

	auditlogdto "github.com/saadahmedbd/Treestore/Rest/DTO/AuditLogDTO"
)

// Get activity timeline
func (r *auditLogRepository) GetActivityTimeline(startDate, endDate time.Time, groupBy string) ([]auditlogdto.ActivityTimelineItem, error) {
	var items []auditlogdto.ActivityTimelineItem

	dateTrunc := "day"
	if groupBy == "hour" {
		dateTrunc = "hour"
	} else if groupBy == "week" {
		dateTrunc = "week"
	} else if groupBy == "month" {
		dateTrunc = "month"
	}

	query := fmt.Sprintf(`
		SELECT 
			DATE_TRUNC('%s', created_at) as date,
			COUNT(*) as total_count
		FROM audit_logs
		WHERE created_at BETWEEN ? AND ?
		GROUP BY DATE_TRUNC('%s', created_at)
		ORDER BY date ASC
	`, dateTrunc, dateTrunc)

	rows, err := r.db.Raw(query, startDate, endDate).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item auditlogdto.ActivityTimelineItem
		var date time.Time
		rows.Scan(&date, &item.TotalCount)
		item.Date = date.Format("2006-01-02 15:04")
		items = append(items, item)
	}

	return items, nil
}
