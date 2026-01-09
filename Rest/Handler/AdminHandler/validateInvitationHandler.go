package adminhandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Adminhandler) ValidateInvitation(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Token is required", http.StatusBadRequest)
		return
	}
	invitation, err := h.adminService.ValidateInvitationToken(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.SendData(w, invitation, 200)
}
