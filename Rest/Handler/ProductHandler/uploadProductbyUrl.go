package producthandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	productservice "github.com/saadahmedbd/Treestore/Rest/Service/ProductService"
	util "github.com/saadahmedbd/Treestore/Util"
)

// ============= UPLOAD VIA JSON (URL-based) =============
func (h *Handler) AddProductImageByURL(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, `{"error":"only sellers can add product images"}`, http.StatusForbidden)
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, `{"error":"invalid user_id format"}`, http.StatusBadRequest)
		return
	}

	// Parse JSON request
	var req productservice.AddImageByURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.ImageURL == "" {
		http.Error(w, `{"error":"image_url is required"}`, http.StatusBadRequest)
		return
	}

	// Add image by URL
	imageData, err := h.service.AddProductImageByURL(uint(id), uint(userID), req)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	util.SendData(w, imageData, http.StatusCreated)
}
