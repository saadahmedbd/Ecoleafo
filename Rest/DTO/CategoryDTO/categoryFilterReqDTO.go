package categorydto

// CategoryFilterRequest - Query parameters for filtering categories
type CategoryFilterRequest struct {
	ParentID   *uint  `json:"parent_id"`   // Filter by parent
	IsFeatured *bool  `json:"is_featured"` // Featured only
	IsActive   *bool  `json:"is_active"`   // Active/inactive
	Search     string `json:"search"`      // Search by name
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
}
