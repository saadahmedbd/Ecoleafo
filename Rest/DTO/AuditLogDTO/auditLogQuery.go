package auditlogdto

// Query parameters for filtering logs
type AuditLogQuery struct {
	Page        int    `json:"page"`
	Limit       int    `json:"limit"`
	ActorID     *uint  `json:"actor_id"`
	ActorType   string `json:"actor_type"`
	Action      string `json:"action"`
	ActionGroup string `json:"action_group"`
	EntityType  string `json:"entity_type"`
	EntityID    *uint  `json:"entity_id"`
	Status      string `json:"status"`
	Severity    string `json:"severity"`
	Category    string `json:"category"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Search      string `json:"search"`
	IPAddress   string `json:"ip_address"`
}
