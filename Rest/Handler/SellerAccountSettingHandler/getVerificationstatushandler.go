package selleraccountsettinghandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// ==========================================
// GET VERIFICATION STATUS
// GET /api/seller/verification/status
// ==========================================

func (h *Selleraccountsettinghandler) GetVerificationStatus(w http.ResponseWriter, r *http.Request) {
	sellerID := r.Context().Value("seller_id").(uint)

	status, err := h.service.GetVerificationStatus(sellerID)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, status, "")
}
