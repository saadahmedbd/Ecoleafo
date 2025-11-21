package commissionpayoutearningshandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GET /api/admin/sellers/earnings/all?page=1&limit=20
func (h *CommissionHandler) GetAllSellerEarnings(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 20
	}

	earnings, total, err := h.commissionService.GetAllSellerEarnings(page, limit)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to get earnings")
		return
	}
	util.SendPaginatedresponse(w, http.StatusOK, "Payout history retrieved successfully", earnings, page, limit, total)

}
