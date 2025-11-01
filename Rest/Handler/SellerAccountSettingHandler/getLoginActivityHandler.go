package selleraccountsettinghandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// ==========================================
// GET LOGIN ACTIVITY
// GET /api/seller/security/activity
// ==========================================

func (h *Selleraccountsettinghandler) GetLoginActivity(w http.ResponseWriter, r *http.Request) {
	sellerID := r.Context().Value("seller_id").(uint)

	activities, err := h.service.GetLoginActivity(sellerID)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, activities, "")
}
