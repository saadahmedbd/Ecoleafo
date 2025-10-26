package categoryHandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GetFeaturedCategories handles GET requests for featured categories
// GET /api/categories/featured?limit=10
func (h *CategoryHandler) GetFeaturedCategories(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit <= 0 {
		limit = 10 // Default limit
	}

	categories, err := h.categoryService.GetFeaturedCategories(limit)
	if err != nil {
		util.RespondJSON(w, http.StatusInternalServerError, nil, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, categories, "Featured categories retrieved successfully")
}
