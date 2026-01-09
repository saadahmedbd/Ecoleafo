package guestcarthandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	address "github.com/saadahmedbd/Treestore/Rest/DTO/Address"
	util "github.com/saadahmedbd/Treestore/Util"
)

// POST /api/checkout/address (Protected)
func (h *GuestcartHandler) CreateCheckoutAddress(w http.ResponseWriter, r *http.Request) {
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

	var req address.CreateAddressRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}

	address, err := h.service.CreateCheckoutAddress(uint(userID), req)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	util.SendData(w, address, http.StatusCreated)
}
