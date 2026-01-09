package producthandler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	// Extract product ID from URL
	productId := r.PathValue("productId")
	Pid, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Get user info from JWT (set by middleware)
	userIDStr := r.Header.Get("user_id")
	userType := r.Header.Get("user_role")

	// Debug logging
	fmt.Printf("DEBUG Delete - user_id: %s, user_role: %s, product_id: %d\n", userIDStr, userType, Pid)

	// Validation
	if userIDStr == "" {
		http.Error(w, `{"error":"user_id not found in headers"}`, http.StatusUnauthorized)
		return
	}
	if userType == "" {
		http.Error(w, `{"error":"user_role not found in headers"}`, http.StatusUnauthorized)
		return
	}

	// Check if user is a seller or admin
	if !strings.Contains(userType, "seller") && !strings.Contains(userType, "admin") {
		http.Error(w, `{"error":"only sellers can delete their products"}`, http.StatusForbidden)
		return
	}

	// Parse user ID from JWT
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, `{"error":"invalid user_id format"}`, http.StatusBadRequest)
		return
	}

	fmt.Printf("DEBUG: About to delete product %d for user %d\n", uint(Pid), uint(userID))

	// Delete product
	err = h.service.DeleteProduct(uint(Pid), uint(userID), strings.Contains(userType, "admin"))
	if err != nil {
		fmt.Printf("DEBUG: Delete error: %v\n", err)
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	// Return success response
	response := map[string]interface{}{
		"message":    "Product deleted successfully",
		"product_id": Pid,
	}
	util.SendData(w, response, http.StatusOK)
}
