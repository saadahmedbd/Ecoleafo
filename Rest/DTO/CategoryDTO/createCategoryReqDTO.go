package categorydto

// CreateCategoryRequest - Request payload for creating a new category
type CreateCategoryRequest struct {
	Name            string `json:"name" validate:"required,min=2,max=100"`
	Description     string `json:"description"`
	Image           string `json:"image" validate:"omitempty,url"`
	Icon            string `json:"icon"`
	ParentID        *uint  `json:"parent_id"` // Null for root category
	SortOrder       int    `json:"sort_order"`
	IsFeatured      bool   `json:"is_featured"`
	IsActive        bool   `json:"is_active"`
	MetaTitle       string `json:"meta_title"`
	MetaDescription string `json:"meta_description"`
	MetaKeywords    string `json:"meta_keywords"`
}
