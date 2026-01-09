package producthandler

import (
	"encoding/json"
	"strings"

	"net/http"
	"strconv"

	productservice "github.com/saadahmedbd/Treestore/Rest/Service/ProductService"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("user_id")
	userType := r.Header.Get("user_role")

	if userIDStr == "" {
		http.Error(w, "user_id not found in headers", http.StatusUnauthorized)
		return
	}
	if userType == "" {
		http.Error(w, "user_role not found in headers", http.StatusUnauthorized)
		return
	}
	// Check if user is a seller - adjust based on your role format
	// Since roles is []string in JWT, it comes as "[seller]" or "[admin seller]"
	if !strings.Contains(userType, "seller") {
		http.Error(w, `{"error":"only sellers can create products"}`, http.StatusForbidden)
		return
	}
	// Parse user ID (this is User.ID from JWT)
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "invalid user_id format", http.StatusBadRequest)
		return
	}
	var req productservice.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	// Create product using User.ID
	product, err := h.service.CreateProduct(req, uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.SendData(w, product, http.StatusCreated)

}
