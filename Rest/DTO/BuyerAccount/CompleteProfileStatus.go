package buyeraccount

type ProfileCompletionStatus struct {
	IsProfileComplete bool     `json:"is_profile_complete"`
	HasPhone          bool     `json:"has_phone"`
	HasAddress        bool     `json:"has_address"`
	MissingFields     []string `json:"missing_fields"`
	NextStep          string   `json:"next_step"` // "complete_profile", "add_address", "ready_to_checkout"
}
