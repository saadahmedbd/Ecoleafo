package categorydto

// UpdateCategoryRequest - Request payload for updating existing category
type UpdateCategoryRequest struct {
	Name            *string `json:"name" validate:"omitempty,min=2,max=100"`
	Description     *string `json:"description"`
	ParentID        *uint   `json:"parent_id"`
	SortOrder       *int    `json:"sort_order"`
	IsFeatured      *bool   `json:"is_featured"`
	IsActive        *bool   `json:"is_active"`
	MetaTitle       *string `json:"meta_title"`
	MetaDescription *string `json:"meta_description"`
	MetaKeywords    *string `json:"meta_keywords"`
}
