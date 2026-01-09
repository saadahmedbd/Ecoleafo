package categoryHandler

import (
	"net/http"
	"strconv"

	categorydto "github.com/saadahmedbd/Treestore/Rest/DTO/CategoryDTO"
	util "github.com/saadahmedbd/Treestore/Util"
)

// GetAllCategories handles GET requests to retrieve all categories with filters
// GET /api/categories?page=1&limit=20&is_featured=true&search=phone
func (h *CategoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	search := r.URL.Query().Get("search")

	// Parse boolean filters
	var isFeatured, isActive *bool
	if featuredStr := r.URL.Query().Get("is_featured"); featuredStr != "" {
		featured := featuredStr == "true"
		isFeatured = &featured
	}
	if activeStr := r.URL.Query().Get("is_active"); activeStr != "" {
		active := activeStr == "true"
		isActive = &active
	}

	// Parse parent ID filter
	var parentID *uint
	if parentIDStr := r.URL.Query().Get("parent_id"); parentIDStr != "" {
		if parentIDStr == "null" {
			// Request for root categories
			parentID = nil
		} else {
			pid, err := strconv.ParseUint(parentIDStr, 10, 32)
			if err == nil {
				pidUint := uint(pid)
				parentID = &pidUint
			}
		}
	}

	// Build filter
	filter := categorydto.CategoryFilterRequest{
		Page:       page,
		Limit:      limit,
		Search:     search,
		IsFeatured: isFeatured,
		IsActive:   isActive,
		ParentID:   parentID,
	}

	// Get categories
	categories, total, err := h.categoryService.GetAllCategories(filter)
	if err != nil {
		util.RespondJSON(w, http.StatusInternalServerError, nil, err.Error())
		return
	}

	// Build response with pagination metadata
	response := map[string]interface{}{
		"categories": categories,
		"pagination": map[string]interface{}{
			"total":       total,
			"page":        filter.Page,
			"limit":       filter.Limit,
			"total_pages": (total + int64(filter.Limit) - 1) / int64(filter.Limit),
		},
	}

	util.RespondJSON(w, http.StatusOK, response, "Categories retrieved successfully")
}
