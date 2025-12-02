package adminhandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Adminhandler) DeactivateAdmin(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("user_id")

	if userIDStr == "" {
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

	err = h.adminService.DeactivateAdmin(uint(adminID), uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	util.SendData(w, "admin deactivate successful", 200)
}
