package categoryHandler

import (
	"encoding/json"
	"strings"

	"net/http"
	"strconv"

	categorydto "github.com/saadahmedbd/Treestore/Rest/DTO/CategoryDTO"
	util "github.com/saadahmedbd/Treestore/Util"
)

// UpdateCategory handles PUT requests to update a category
// PUT /api/categories/update?id=1
func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
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

	// Parse request body
	var req categorydto.UpdateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.RespondJSON(w, http.StatusBadRequest, nil, "Invalid request body")
		return
	}

	// Update category
	category, err := h.categoryService.UpdateCategory(uint(id), req)
	if err != nil {
		util.RespondJSON(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, category, "Category updated successfully")
}
