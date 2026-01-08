package categoryHandler

import (
	"net/http"
	"strconv"
	"strings"

	util "github.com/saadahmedbd/Treestore/Util"
)

// UploadCategoryImage handles POST requests to upload category image
// POST /api/categories/upload-image?id=1
func (h *CategoryHandler) UploadCategoryImage(w http.ResponseWriter, r *http.Request) {
	// Check admin role
	userType := r.Header.Get("user_role")
	if !strings.Contains(userType, "admin") {
		http.Error(w, `{"error":"only admin can upload category images"}`, http.StatusForbidden)
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
	if err := r.ParseMultipartForm(5 << 20); err != nil {
		util.RespondJSON(w, http.StatusBadRequest, nil, "Failed to parse form")
		return
	}

	// Get uploaded file
	file, header, err := r.FormFile("image")
	if err != nil {
		util.RespondJSON(w, http.StatusBadRequest, nil, "Image file is required")
		return
	}
	defer file.Close()

	// Upload to Cloudinary
	cfg := util.DefaultImageConfig("categories")

	imageURL, err := util.UploadImageToCloudinary(file, header, cfg)
	if err != nil {
		util.RespondJSON(w, http.StatusInternalServerError, nil, err.Error())
		return
	}

	// Update category with image URL
	category, err := h.categoryService.UpdateCategoryImage(uint(id), imageURL)
	if err != nil {
		util.RespondJSON(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, category, "Category image uploaded successfully")
}
