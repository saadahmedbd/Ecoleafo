package categorydto

// CategoryBriefResponse - Lightweight response for nested categories
type CategoryBriefResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	Image      string `json:"image"`
	Icon       string `json:"icon"`
	IsActive   bool   `json:"is_active"`
	IsFeatured bool   `json:"is_featured"`
}
