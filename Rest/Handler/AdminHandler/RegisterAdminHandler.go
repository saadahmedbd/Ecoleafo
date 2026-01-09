package adminhandler

import (
	"encoding/json"
	"net/http"

	admin "github.com/saadahmedbd/Treestore/Rest/DTO/Admin"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Adminhandler) RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	var req admin.RegisterAdminRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	admin, err := h.adminService.RegisterAdmin(req)
	if err != nil {
		http.Error(w, "Failed to register admin", http.StatusInternalServerError)
		return
	}
	util.SendData(w, admin, 200)
}
