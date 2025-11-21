package commissionpayoutearningshandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// Get /api/admin/commission/settings

func (h *CommissionHandler) GetCommissionSettings(w http.ResponseWriter, r *http.Request) {
	settings, err := h.commissionService.GetCommissionSettings()
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to get settings")
		return
	}

	util.RespondJSON(w, http.StatusOK, settings, "Commission settings retrieved successfully")
}
