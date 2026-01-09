package categoryHandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GetCategoryByID handles GET requests to retrieve a category by ID
// GET /api/categories/get?id=1
func (h *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	// Parse category ID from query parameter
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		util.RespondJSON(w, http.StatusBadRequest, nil, "Category ID is required")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		util.RespondJSON(w, http.StatusBadRequest, nil, "Invalid category ID")
		return
	}

	// Get category
	category, err := h.categoryService.GetCategoryByID(uint(id))
	if err != nil {
		util.RespondJSON(w, http.StatusNotFound, nil, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, category, "Category retrieved successfully")
}
