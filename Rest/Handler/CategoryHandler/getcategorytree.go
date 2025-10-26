package categoryHandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GetCategoryTree handles GET requests for hierarchical category tree
// GET /api/categories/tree
func (h *CategoryHandler) GetCategoryTree(w http.ResponseWriter, r *http.Request) {
	tree, err := h.categoryService.GetCategoryTree()
	if err != nil {
		util.RespondJSON(w, http.StatusInternalServerError, nil, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, tree, "Category tree retrieved successfully")
}
