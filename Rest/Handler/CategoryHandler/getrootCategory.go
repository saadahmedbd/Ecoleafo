package categoryHandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GetRootCategories handles GET requests for top-level categories
// GET /api/categories/root
func (h *CategoryHandler) GetRootCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.categoryService.GetRootCategories()
	if err != nil {
		util.RespondJSON(w, http.StatusInternalServerError, nil, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, categories, "Root categories retrieved successfully")
}
