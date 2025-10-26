package categoryHandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GetBreadcrumb handles GET requests for category breadcrumb
// GET /api/categories/breadcrumb?id=5
func (h *CategoryHandler) GetBreadcrumb(w http.ResponseWriter, r *http.Request) {
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

	breadcrumb, err := h.categoryService.GetBreadcrumb(uint(id))
	if err != nil {
		util.RespondJSON(w, http.StatusNotFound, nil, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, breadcrumb, "Breadcrumb retrieved successfully")
}
