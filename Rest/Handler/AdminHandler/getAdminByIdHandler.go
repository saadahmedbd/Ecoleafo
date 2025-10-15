package adminhandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Adminhandler) GetAdminByID(w http.ResponseWriter, r *http.Request) {

	userType := r.Header.Get("user_role")

	if userType != "[admin]" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	adminID, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		http.Error(w, "invalid admin id", http.StatusUnauthorized)
		return
	}
	admin, err := h.adminService.GetAdminByID(uint(adminID))
	if err != nil {
		http.Error(w, "admin not found", http.StatusUnauthorized)
		return
	}
	util.SendData(w, admin, 200)

}
