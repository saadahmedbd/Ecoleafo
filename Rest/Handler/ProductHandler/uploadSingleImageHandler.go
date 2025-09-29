package producthandler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Handler) UploadImage(w http.ResponseWriter, r *http.Request) {
	ImageId := r.PathValue("ImageId")
	id, err := strconv.Atoi(ImageId)
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

	// Parse multipart form (10MB max)
	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, `{"error":"failed to parse form data"}`, http.StatusBadRequest)
		return
	}

	// Get the image file
	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		http.Error(w, `{"error":"image file is required"}`, http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Get optional form fields
	altText := r.FormValue("alt_text")
	isPrimaryStr := r.FormValue("is_primary")
	sortOrderStr := r.FormValue("sort_order")

	// Parse boolean and int fields
	isPrimary := isPrimaryStr == "true"
	sortOrder := 1
	if sortOrderStr != "" {
		if so, err := strconv.Atoi(sortOrderStr); err == nil {
			sortOrder = so
		}
	}

	// Upload the image
	imageData, err := h.service.UploadProductImage(uint(id), uint(userID), file, fileHeader, altText, isPrimary, sortOrder)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	util.SendData(w, imageData, http.StatusCreated)
}
