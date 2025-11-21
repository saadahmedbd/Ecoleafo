package commissionpayoutearningshandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GET /api/admin/sellers/earnings?seller_id=5
func (h *CommissionHandler) GetSellerEarnings(w http.ResponseWriter, r *http.Request) {
	sellerIDStr := r.URL.Query().Get("seller_id")
	if sellerIDStr == "" {
		util.RespondError(w, http.StatusBadRequest, "seller_id is required")
		return
	}

	sellerID, err := strconv.ParseUint(sellerIDStr, 10, 64)
	if err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid seller_id")
		return
	}

	earnings, err := h.commissionService.GetSellerEarnings(uint(sellerID))
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to get earnings")
		return
	}

	util.RespondJSON(w, http.StatusOK, earnings, "Seller earnings retrieved successfully")
}
