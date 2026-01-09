package buyerprofilehandler

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Buyerprofilehandler) GetBuyerProfile(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("user_id")
	userType := r.Header.Get("user_role")
	log.Println(userType)

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
	profile, err := h.buyerservice.GetBuyerProfile(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.SendData(w, profile, 200)

}
