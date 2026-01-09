package buyerprofilehandler

import (
	"net/http"
	"strconv"
	"strings"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Buyerprofilehandler) DeleteAddress(w http.ResponseWriter, r *http.Request) {
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
	addressIDstr := r.PathValue("id")
	addressID, err := strconv.Atoi(addressIDstr)
	if err != nil {
		http.Error(w, "id not convert", http.StatusNoContent)
		return
	}

	err = h.buyerservice.DeleteAddress(uint(userID), uint(addressID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{
		"message": "Address deleted successfully",
	}
	util.SendData(w, response, http.StatusOK)

}
