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

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	//get seller if from jwt
	userIDStr := r.Header.Get("user_id")
	userType := r.Header.Get("user_role")

	productId := r.PathValue("productId")
	sId, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "id not convert", http.StatusNoContent)
		return
	}

	// Validation
	if userIDStr == "" {
		http.Error(w, `{"error":"user_id not found in headers"}`, http.StatusUnauthorized)
		return
	}
	if userType == "" {
		http.Error(w, `{"error":"user_role not found in headers"}`, http.StatusUnauthorized)
		return
	}

	// Check if user is a seller
	if !strings.Contains(userType, "seller") {
		http.Error(w, `{"error":"only sellers can update products"}`, http.StatusForbidden)
		return
	}
	sellerID, err := strconv.ParseInt(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var req productservice.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return

	}

	product, err := h.service.UpdateProduct(uint(sId), req, uint(sellerID))
	if err != nil {
		if strings.Contains(err.Error(), "product not found") || strings.Contains(err.Error(), "access denied") {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Failed to update product: %v", err), http.StatusInternalServerError)
		return
	}

	util.SendData(w, product, 200)

}
