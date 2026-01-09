package adminhandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Adminhandler) GetAdminProfile(w http.ResponseWriter, r *http.Request) {
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
	admin, err := h.adminService.GetAdminByUserID(uint(userID))
	if err != nil {
		http.Error(w, "failed to get admin profile", http.StatusInternalServerError)
		return
	}
	util.SendData(w, admin, 200)
}
