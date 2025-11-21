package commissionpayoutearningshandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GET /api/admin/earnings/overview
func (h *CommissionHandler) GetPlatformEarningsOverview(w http.ResponseWriter, r *http.Request) {
	overview, err := h.commissionService.GetPlatformEarningsOverview()
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to get earnings overview")
		return
	}

	util.RespondJSON(w, http.StatusOK, overview, "Platform earnings overview retrieved successfully")
}

// GET /api/admin/revenue/monthly?year=2025
func (h *CommissionHandler) GetMonthlyRevenue(w http.ResponseWriter, r *http.Request) {
	year, _ := strconv.Atoi(r.URL.Query().Get("year"))

	revenue, err := h.commissionService.GetMonthlyRevenue(year)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to get monthly revenue")
		return
	}

	util.RespondJSON(w, http.StatusOK, revenue, "Monthly revenue retrieved successfully")
}

// GET /api/admin/sellers/top?limit=10
func (h *CommissionHandler) GetTopSellers(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit < 1 {
		limit = 10
	}

	sellers, err := h.commissionService.GetTopSellers(limit)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to get top sellers")
		return
	}

	util.RespondJSON(w, http.StatusOK, sellers, "Top sellers retrieved successfully")
}
