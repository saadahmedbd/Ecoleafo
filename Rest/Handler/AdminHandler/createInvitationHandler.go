package adminhandler

import (
	"encoding/json"

	"log"
	"net/http"
	"strconv"

	admin "github.com/saadahmedbd/Treestore/Rest/DTO/Admin"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Adminhandler) CreateInvitation(w http.ResponseWriter, r *http.Request) {
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
	log.Println("FindByUserID called with:", userID)

	// Get admin ID from user ID
	admins, err := h.adminService.GetAdminByUserID(uint(userID))
	if err != nil {
		http.Error(w, "admin not found", http.StatusNotFound)
		return
	}

	var req admin.CreateAdminInvitationRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusForbidden)
		return
	}
	invitation, err := h.adminService.CreateInvitation(admins.UserID, req)
	if err != nil {
		http.Error(w, "failed to create invitation", http.StatusInternalServerError)
		return
	}
	util.SendData(w, invitation, 200)
}
