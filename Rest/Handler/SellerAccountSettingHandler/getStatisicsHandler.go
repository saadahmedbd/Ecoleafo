package selleraccountsettinghandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// ==========================================
// GET STATISTICS
// GET /api/seller/statistics
// ==========================================

func (h *Selleraccountsettinghandler) GetStatistics(w http.ResponseWriter, r *http.Request) {
	sellerID := r.Context().Value("seller_id").(uint)

	stats, err := h.service.GetStatistics(sellerID)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, stats, "")
}
