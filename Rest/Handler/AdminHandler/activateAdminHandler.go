package adminhandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Adminhandler) ActivateAdmin(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("user_id")
	userType := r.Header.Get("user_role")
	if userIDStr == "" || userType == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if userType != "[admin]" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	//convert useridstr
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusForbidden)
		return
	}

	adminID, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		http.Error(w, "invalid admin id", http.StatusBadRequest)
		return
	}

	err = h.adminService.ActivateAdmin(uint(adminID), uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	util.SendData(w, "admin Activated successful", 200)
}
