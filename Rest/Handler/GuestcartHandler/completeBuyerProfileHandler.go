package guestcarthandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	buyeraccount "github.com/saadahmedbd/Treestore/Rest/DTO/BuyerAccount"
	util "github.com/saadahmedbd/Treestore/Util"
)

// POST /api/checkout/complete-profile (Protected)
func (h *GuestcartHandler) CompleteProfile(w http.ResponseWriter, r *http.Request) {
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

	var req buyeraccount.CompleteProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}

	if err := h.service.CompleteProfile(uint(userID), req); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Profile completed successfully",
	}
	util.SendData(w, response, http.StatusOK)
}
