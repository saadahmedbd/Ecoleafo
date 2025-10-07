package selleraccounthandler

import (
	"net/http"
	"strconv"
	"strings"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GET /api/seller/profile (Protected)
func (h *SellerRegistrationHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("user_id")
	userType := r.Header.Get("user_role")

	if userIDStr == "" || userType == "" {
		http.Error(w, `{"error":"authentication required"}`, http.StatusUnauthorized)
		return
	}

	if !strings.Contains(userType, "seller") {
		http.Error(w, `{"error":"only sellers can access this"}`, http.StatusForbidden)
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, `{"error":"invalid user_id"}`, http.StatusBadRequest)
		return
	}
	profile, err := h.service.GetSellerProfile(uint(userID))
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	util.SendData(w, profile, http.StatusOK)

}
