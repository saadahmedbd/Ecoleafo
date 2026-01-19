package categoryHandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GetCategoriesForSeller returns all categories without is_active filter for product creation
func (h *CategoryHandler) GetCategoriesForSeller(w http.ResponseWriter, r *http.Request) {
	categories, err := h.categoryService.GetCategoriesForSeller()
	if err != nil {
		util.RespondJSON(w, http.StatusInternalServerError, nil, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, map[string]interface{}{"categories": categories}, "Categories retrieved successfully")
}
