package selleraccountsettinghandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// ==========================================
// GET FULL PROFILE
// GET /api/seller/profile
// ==========================================

func (h *Selleraccountsettinghandler) GetFullProfile(w http.ResponseWriter, r *http.Request) {
	// Get seller ID from context (set by auth middleware)
	sellerID := r.Context().Value("seller_id").(uint)

	profile, err := h.service.GetFullProfile(sellerID)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, profile, "success")
}
