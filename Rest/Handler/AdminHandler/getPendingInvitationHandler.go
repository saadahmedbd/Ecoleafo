package adminhandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Adminhandler) GetPendingInvitations(w http.ResponseWriter, r *http.Request) {

	invitations, err := h.adminService.GetPendingInvitations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, invitations, http.StatusOK)
}
