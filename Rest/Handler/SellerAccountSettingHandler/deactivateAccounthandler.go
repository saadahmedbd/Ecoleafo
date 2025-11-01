package selleraccountsettinghandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// ==========================================
// DEACTIVATE ACCOUNT
// POST /api/seller/account/deactivate
// ==========================================

func (h *Selleraccountsettinghandler) DeactivateAccount(w http.ResponseWriter, r *http.Request) {
	sellerID := r.Context().Value("seller_id").(uint)

	response, err := h.service.DeactivateAccount(sellerID)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, response, "deactivate Account Seccesfully")
}
