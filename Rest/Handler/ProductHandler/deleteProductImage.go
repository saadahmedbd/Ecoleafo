package producthandler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	util "github.com/saadahmedbd/Treestore/Util"
)

// ============= DELETE PRODUCT IMAGE =============
func (h *Handler) DeleteProductImage(w http.ResponseWriter, r *http.Request) {
	// Extract IDs from URL
	productIDStr := r.PathValue("productId")
	imageIDStr := r.PathValue("imageId")

	if productIDStr == "" || imageIDStr == "" {
		http.Error(w, `{"error":"product ID and image ID are required"}`, http.StatusBadRequest)
		return
	}

	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		http.Error(w, `{"error":"invalid product ID"}`, http.StatusBadRequest)
		return
	}

	imageID, err := strconv.ParseUint(imageIDStr, 10, 32)
	if err != nil {
		http.Error(w, `{"error":"invalid image ID"}`, http.StatusBadRequest)
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
		http.Error(w, `{"error":"only sellers can delete product images"}`, http.StatusForbidden)
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, `{"error":"invalid user_id format"}`, http.StatusBadRequest)
		return
	}

	// Delete image
	err = h.service.DeleteProductImage(uint(productID), uint(imageID), uint(userID))
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message":    "Image deleted successfully",
		"product_id": productID,
		"image_id":   imageID,
	}

	util.SendData(w, response, http.StatusOK)
}
