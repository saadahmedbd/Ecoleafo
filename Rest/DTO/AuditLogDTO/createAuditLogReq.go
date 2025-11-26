package auditlogdto

// Create audit log request (for manual logging)
type CreateAuditLogRequest struct {
	Action      string                 `json:"action" binding:"required"`
	Description string                 `json:"description"`
	EntityType  string                 `json:"entity_type"`
	EntityID    *uint                  `json:"entity_id"`
	EntityName  string                 `json:"entity_name"`
	OldValues   map[string]interface{} `json:"old_values"`
	NewValues   map[string]interface{} `json:"new_values"`
	Status      string                 `json:"status"`
	Severity    string                 `json:"severity"`
	Category    string                 `json:"category"`
	Metadata    map[string]interface{} `json:"metadata"`
}
