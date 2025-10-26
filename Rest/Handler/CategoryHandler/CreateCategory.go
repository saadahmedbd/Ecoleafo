package categoryHandler

import (
	"encoding/json"
	"net/http"
	"strings"

	categorydto "github.com/saadahmedbd/Treestore/Rest/DTO/CategoryDTO"
	util "github.com/saadahmedbd/Treestore/Util"
)

// CreateCategory handles POST requests to create a new category
// POST /api/categories
func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
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
	var req categorydto.CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.RespondJSON(w, http.StatusBadRequest, nil, "Invalid request body")
		return
	}

	// Validate request
	if err := util.ValidateStruct(req); err != nil {
		util.RespondJSON(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	// Create category
	category, err := h.categoryService.CreateCategory(req)
	if err != nil {
		util.RespondJSON(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusCreated, category, "Category created successfully")
}
