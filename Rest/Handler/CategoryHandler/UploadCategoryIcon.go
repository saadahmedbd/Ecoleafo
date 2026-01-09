package categoryHandler

import (
	"net/http"
	"strconv"
	"strings"

	util "github.com/saadahmedbd/Treestore/Util"
)

// UploadCategoryIcon handles POST requests to upload category icon
// POST /api/categories/upload-icon?id=1
func (h *CategoryHandler) UploadCategoryIcon(w http.ResponseWriter, r *http.Request) {
	// Check admin role
	userType := r.Header.Get("user_role")
	if !strings.Contains(userType, "admin") {
		http.Error(w, `{"error":"only admin can upload category icons"}`, http.StatusForbidden)
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

	// Parse multipart form
	if err := r.ParseMultipartForm(2 << 20); err != nil {
		util.RespondJSON(w, http.StatusBadRequest, nil, "Failed to parse form")
		return
	}

	// Get uploaded file
	file, header, err := r.FormFile("icon")
	if err != nil {
		util.RespondJSON(w, http.StatusBadRequest, nil, "Icon file is required")
		return
	}
	defer file.Close()

	// Upload to Cloudinary
	cfg := util.DefaultImageConfig("category-icons")

	iconURL, err := util.UploadImageToCloudinary(file, header, cfg)
	if err != nil {
		util.RespondJSON(w, http.StatusInternalServerError, nil, err.Error())
		return
	}

	// Update category with icon URL
	category, err := h.categoryService.UpdateCategoryIcon(uint(id), iconURL)
	if err != nil {
		util.RespondJSON(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, category, "Category icon uploaded successfully")
}
