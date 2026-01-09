package categorydto

// CategoryTreeResponse - Hierarchical tree structure for navigation menus
type CategoryTreeResponse struct {
	ID       uint                   `json:"id"`
	Name     string                 `json:"name"`
	Slug     string                 `json:"slug"`
	Icon     string                 `json:"icon"`
	Children []CategoryTreeResponse `json:"children,omitempty"`
}
