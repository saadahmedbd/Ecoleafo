package auditlogdto

type TopActorItem struct {
	ActorID     uint   `json:"actor_id"`
	ActorName   string `json:"actor_name"`
	ActorType   string `json:"actor_type"`
	ActionCount int    `json:"action_count"`
}
