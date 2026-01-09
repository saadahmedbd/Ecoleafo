package commissionpayoutearningshandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GET /api/admin/payouts/pending?page=1&limit=20
func (h *CommissionHandler) GetPendingPayouts(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 20
	}

	payouts, total, err := h.commissionService.GetPendingPayouts(page, limit)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to get payouts")
		return
	}

	util.SendPaginatedresponse(w, http.StatusOK, "Payout history retrieved successfully", payouts, page, limit, total)

}
