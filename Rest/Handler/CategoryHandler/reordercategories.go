package categoryHandler

import (
	"encoding/json"
	"net/http"
	"strings"

	util "github.com/saadahmedbd/Treestore/Util"
)

// ReorderCategories handles POST requests to update category sort order
// POST /api/categories/reorder
func (h *CategoryHandler) ReorderCategories(w http.ResponseWriter, r *http.Request) {
	// Check if user has admin role
	userType := r.Header.Get("user_role")
	if userType == "" {
		http.Error(w, "user_role not found in headers", http.StatusUnauthorized)
		return
	}
	// Check if user is a seller - adjust based on your role format
	// Since roles is []string in JWT, it comes as "[seller]" or "[admin seller]"
	if !strings.Contains(userType, "admin") {
		http.Error(w, `{"error":"only admin can create products"}`, http.StatusForbidden)
		return
	}

	// Parse request body
	var updatess []struct {
		ID        uint `json:"id"`
		SortOrder int  `json:"sort_order"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updatess); err != nil {
		util.RespondJSON(w, http.StatusBadRequest, nil, "Invalid request body")
		return
	}

	// Update sort order
	if err := h.categoryService.ReorderCategories([]struct {
		ID        uint
		SortOrder int
	}(updatess)); err != nil {
		util.RespondJSON(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, nil, "Categories reordered successfully")
}
