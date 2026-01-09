package commissionpayoutearningshandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GET /api/admin/payouts/history?seller_id=5&page=1&limit=20
func (h *CommissionHandler) GetPayoutHistory(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 20
	}

	var sellerID *uint
	sellerIDStr := r.URL.Query().Get("seller_id")
	if sellerIDStr != "" {
		id, err := strconv.ParseUint(sellerIDStr, 10, 64)
		if err == nil {
			uid := uint(id)
			sellerID = &uid
		}
	}

	payouts, total, err := h.commissionService.GetPayoutHistory(sellerID, page, limit)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to get payout history")
		return
	}
	util.SendPaginatedresponse(w, http.StatusOK, "Payout history retrieved successfully", payouts, page, limit, total)

}
