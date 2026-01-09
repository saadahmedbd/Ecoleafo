package guestcarthandler

import (
	"fmt"
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GET /api/checkout/profile-status (Protected)
func (h *GuestcartHandler) CheckProfileStatus(w http.ResponseWriter, r *http.Request) {
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

	status, err := h.service.CheckProfileCompletion(uint(userID))
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	util.SendData(w, status, http.StatusOK)
}
