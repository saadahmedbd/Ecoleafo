package sellerproflehandler

import (
	"net/http"
	"strconv"
	"strings"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *SellerProfileHandler) GetSellerProfile(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("user_id")
	userType := r.Header.Get("user_role")
	if userIDStr == "" || userType == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if !strings.Contains(userType, "seller") {
		http.Error(w, "only seller can acces this account", http.StatusForbidden)
		return
	}
	//comnver user id uint
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	//get seller profile
	profile, err := h.service.GetSellerProfile(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.SendData(w, profile,200)

}
