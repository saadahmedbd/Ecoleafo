package auditlogdto

// Activity timeline item
type ActivityTimelineItem struct {
	Date       string         `json:"date"`
	TotalCount int            `json:"total_count"`
	ByAction   map[string]int `json:"by_action"`
}
