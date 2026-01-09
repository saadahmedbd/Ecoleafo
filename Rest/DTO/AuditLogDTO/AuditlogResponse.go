package auditlogdto

// Response for single audit log
type AuditLogResponse struct {
	ID          uint                   `json:"id"`
	ActorID     *uint                  `json:"actor_id"`
	ActorType   string                 `json:"actor_type"`
	ActorName   string                 `json:"actor_name"`
	ActorEmail  string                 `json:"actor_email"`
	Action      string                 `json:"action"`
	ActionGroup string                 `json:"action_group"`
	ActionLabel string                 `json:"action_label"` // Human readable
	Description string                 `json:"description"`
	EntityType  string                 `json:"entity_type"`
	EntityID    *uint                  `json:"entity_id"`
	EntityName  string                 `json:"entity_name"`
	OldValues   map[string]interface{} `json:"old_values,omitempty"`
	NewValues   map[string]interface{} `json:"new_values,omitempty"`
	Changes     map[string]interface{} `json:"changes,omitempty"`
	IPAddress   string                 `json:"ip_address"`
	UserAgent   string                 `json:"user_agent"`
	Status      string                 `json:"status"`
	Severity    string                 `json:"severity"`
	Category    string                 `json:"category"`
	CreatedAt   string                 `json:"created_at"`
	TimeAgo     string                 `json:"time_ago"`
}
