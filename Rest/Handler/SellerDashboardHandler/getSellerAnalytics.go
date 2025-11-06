package sellerdashboardhandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

// GetSalesAnalytics - GET /api/seller/analytics/sales?period=week
func (h *DashboardHandler) GetSalesAnalytics(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value(constants.ContextKeyUserID)
	userRoleVal := r.Context().Value(constants.ContextKeyRole)
	if userIDVal == nil || userRoleVal == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	sellerID, ok := userIDVal.(uint)
	if !ok {
		http.Error(w, `{"error":"invalid user_id type"}`, http.StatusInternalServerError)
		return
	}
	// Get period from query params (default: week)
	period := r.URL.Query().Get("period")
	if period == "" {
		period = "week"
	}

	analytics, err := h.DashboardService.GetSalesAnalytics(sellerID, period)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, analytics, "Sales analytics retrieved successfully")
}
