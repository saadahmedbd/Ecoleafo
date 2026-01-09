package guestcarthandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

// PUT /api/cart/items/{product_id}
func (h *GuestcartHandler) UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	sessionID := h.getOrCreateSessionID(r, w)
	userID := h.getUserID(r)

	productId := r.PathValue("product_id")
	productID, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "cannot convert id to int", http.StatusBadRequest)
		return
	}

	var req struct {
		Quantity int `json:"quantity"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateCartItemQuantity(sessionID, userID, uint(productID), req.Quantity); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message": "Cart updated successfully",
	}
	util.SendData(w, response, http.StatusOK)
}
