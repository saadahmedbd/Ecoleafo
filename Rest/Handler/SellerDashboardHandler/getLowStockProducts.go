package sellerdashboardhandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

// GetLowStockProducts - GET /api/seller/products/low-stock?threshold=10
func (h *DashboardHandler) GetLowStockProducts(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value(constants.ContextKeyUserID)
	userRoleVal := r.Context().Value(constants.ContextKeyRole)
	if userIDVal == nil || userRoleVal == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	regUserID, ok := userIDVal.(uint)
	if !ok {
		http.Error(w, `{"error":"invalid user_id type"}`, http.StatusInternalServerError)
		return
	}

	// Get threshold from query params
	thresholdStr := r.URL.Query().Get("threshold")
	threshold := 10 // default
	if thresholdStr != "" {
		if parsedThreshold, err := strconv.Atoi(thresholdStr); err == nil {
			threshold = parsedThreshold
		}
	}

	products, err := h.DashboardService.GetLowStockProducts(regUserID, threshold)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"products": products,
		"count":    len(products),
	}, "Low stock products retrieved successfully")
}
