package adminmangementhandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *AdminManagementHandler) GetPlatformEarningsOverview(w http.ResponseWriter, r *http.Request) {
	overview, err := h.commissionService.GetPlatformEarningsOverview()
	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, overview, http.StatusOK)
}
