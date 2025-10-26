package categoryHandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GetSubcategories handles GET requests for child categories
// GET /api/categories/subcategories?parent_id=1
func (h *CategoryHandler) GetSubcategories(w http.ResponseWriter, r *http.Request) {
	parentIDStr := r.URL.Query().Get("parent_id")
	if parentIDStr == "" {
		util.RespondJSON(w, http.StatusBadRequest, nil, "Parent ID is required")
		return
	}

	parentID, err := strconv.ParseUint(parentIDStr, 10, 32)
	if err != nil {
		util.RespondJSON(w, http.StatusBadRequest, nil, "Invalid parent ID")
		return
	}

	categories, err := h.categoryService.GetSubcategories(uint(parentID))
	if err != nil {
		util.RespondJSON(w, http.StatusInternalServerError, nil, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, categories, "Subcategories retrieved successfully")
}
