package categorydto

import "time"

// CategoryResponse - Standard response for category data
type CategoryResponse struct {
	ID              uint                    `json:"id"`
	Name            string                  `json:"name"`
	Slug            string                  `json:"slug"`
	Description     string                  `json:"description"`
	Image           string                  `json:"image"`
	Icon            string                  `json:"icon"`
	ParentID        *uint                   `json:"parent_id"`
	SortOrder       int                     `json:"sort_order"`
	IsFeatured      bool                    `json:"is_featured"`
	IsActive        bool                    `json:"is_active"`
	ProductCount    int64                   `json:"product_count"` // Number of products in category
	MetaTitle       string                  `json:"meta_title"`
	MetaDescription string                  `json:"meta_description"`
	MetaKeywords    string                  `json:"meta_keywords"`
	CreatedAt       time.Time               `json:"created_at"`
	UpdatedAt       time.Time               `json:"updated_at"`
	Parent          *CategoryBriefResponse  `json:"parent,omitempty"`
	Children        []CategoryBriefResponse `json:"children,omitempty"`
}
