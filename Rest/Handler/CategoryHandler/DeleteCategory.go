package categoryHandler

import (
	"net/http"
	"strconv"
	"strings"

	util "github.com/saadahmedbd/Treestore/Util"
)

// DeleteCategory handles DELETE requests to delete a category
// DELETE /api/categories/delete?id=1
func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
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

	// Parse category ID
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

	// Delete category
	if err := h.categoryService.DeleteCategory(uint(id)); err != nil {
		util.RespondJSON(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, nil, "Category deleted successfully")
}
