package guestcarthandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	address "github.com/saadahmedbd/Treestore/Rest/DTO/Address"
	util "github.com/saadahmedbd/Treestore/Util"
)

// POST /api/checkout/initiate (Protected)
func (h *GuestcartHandler) InitiateCheckout(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("user_id")
	if userIDStr == "" {
		http.Error(w, `{"error":"authentication required"}`, http.StatusUnauthorized)
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, `{"error":"invalid user_id"}`, http.StatusBadRequest)
		return
	}

	var req address.CheckoutInitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}

	if err := h.service.InitiateCheckout(uint(userID), req); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message":     "Checkout initiated successfully",
		"can_proceed": true,
	}
	util.SendData(w, response, http.StatusOK)
}
