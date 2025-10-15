package adminhandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	admin "github.com/saadahmedbd/Treestore/Rest/DTO/Admin"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Adminhandler) UpdatePermission(w http.ResponseWriter, r *http.Request) {
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

	var req admin.UpdateAdminPermissionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invliad request body", http.StatusBadRequest)
		return
	}

	updatedAdmin, err := h.adminService.UpdateAdminPermissions(uint(adminID), uint(userID), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	util.SendData(w, updatedAdmin, 200)
}
