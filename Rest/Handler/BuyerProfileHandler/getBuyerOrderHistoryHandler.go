package buyerprofilehandler

import (
	"net/http"
	"strconv"
	"strings"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Buyerprofilehandler) GetBuyerOrderHistory(w http.ResponseWriter, r *http.Request) {
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
	// Get pagination params
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 50 {
		limit = 10
	}

	orderHistory, err := h.buyerservice.GetBuyerOrderHistory(uint(userID), page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.SendData(w, orderHistory, 200)
}
