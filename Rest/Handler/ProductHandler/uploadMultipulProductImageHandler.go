package producthandler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	util "github.com/saadahmedbd/Treestore/Util"
)

// ============= UPLOAD MULTIPLE PRODUCT IMAGES =============
func (h *Handler) UploadMultipleProductImages(w http.ResponseWriter, r *http.Request) {
	// Extract product ID from URL
	imageId := r.PathValue("ImageId")
	id, err := strconv.Atoi(imageId)
	if err != nil {
		http.Error(w, "Invalid image ID", http.StatusBadRequest)
		return
	}
	// Get user info from JWT
	userIDStr := r.Header.Get("user_id")
	userType := r.Header.Get("user_role")

	if userIDStr == "" || userType == "" {
		http.Error(w, `{"error":"authentication required"}`, http.StatusUnauthorized)
		return
	}

	if !strings.Contains(userType, "seller") {
		http.Error(w, `{"error":"only sellers can upload product images"}`, http.StatusForbidden)
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, `{"error":"invalid user_id format"}`, http.StatusBadRequest)
		return
	}

	// Parse multipart form (100MB max for multiple files)
	err = r.ParseMultipartForm(100 << 20)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"failed to parse form data: %s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	// Check if MultipartForm was parsed
	if r.MultipartForm == nil {
		http.Error(w, `{"error":"no multipart form data found"}`, http.StatusBadRequest)
		return
	}

	// Get multiple image files
	files := r.MultipartForm.File["images"]
	if len(files) == 0 {
		http.Error(w, `{"error":"at least one image file is required"}`, http.StatusBadRequest)
		return
	}

	// Upload multiple images
	imageList, err := h.service.UploadMultipleProductImages(uint(id), uint(userID), files)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message":     "Images uploaded successfully",
		"images":      imageList,
		"total_count": len(imageList),
	}

	util.SendData(w, response, http.StatusCreated)
}
