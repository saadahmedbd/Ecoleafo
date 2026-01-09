package categoryHandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GetCategoryBySlug handles GET requests to retrieve a category by slug
// GET /api/categories/slug?slug=electronics
func (h *CategoryHandler) GetCategoryBySlug(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Query().Get("slug")
	if slug == "" {
		util.RespondJSON(w, http.StatusBadRequest, nil, "Category slug is required")
		return
	}

	category, err := h.categoryService.GetCategoryBySlug(slug)
	if err != nil {
		util.RespondJSON(w, http.StatusNotFound, nil, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, category, "Category retrieved successfully")
}
