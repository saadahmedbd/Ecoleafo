package buyerprofilehandler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	buyerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/BuyerProfile"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Buyerprofilehandler) UpdateBuyerProfile(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("user_id")
	userType := r.Header.Get("user_role")
	if userIDStr == "" || userType == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if !strings.Contains("[buyer]", userType) {
		http.Error(w, "Forbidden: Access is denied", http.StatusForbidden)
		return
	}
	// Convert userID to uint
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	var req buyerprofile.UpdateBuyerProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	updatedProfile, err := h.buyerservice.UpdateBuyerProfile(uint(userID), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.SendData(w, updatedProfile, 200)
}
