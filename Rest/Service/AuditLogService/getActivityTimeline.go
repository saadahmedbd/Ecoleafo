package auditlogservice

import (
	"time"

	auditlogdto "github.com/saadahmedbd/Treestore/Rest/DTO/AuditLogDTO"
)

// Get activity timeline
func (s *auditLogService) GetActivityTimeline(startDate, endDate string, groupBy string) ([]auditlogdto.ActivityTimelineItem, error) {
	start, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)

	if start.IsZero() {
		start = time.Now().AddDate(0, 0, -7)
	}
	if end.IsZero() {
		end = time.Now()
	}
	if groupBy == "" {
		groupBy = "day"
	}

	return s.auditRepo.GetActivityTimeline(start, end, groupBy)
}
